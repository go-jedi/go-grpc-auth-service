package app

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/go-jedi/auth-service/internal/tracing"

	"github.com/sony/gobreaker"

	"github.com/go-jedi/auth-service/internal/interceptor"
	"github.com/go-jedi/auth-service/internal/logger"
	"github.com/go-jedi/auth-service/internal/utils/rate_limiter"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/go-jedi/platform_common/pkg/closer"

	"github.com/go-jedi/auth-service/internal/config"
	descAuth "github.com/go-jedi/auth-service/pkg/auth_v1"
	descUser "github.com/go-jedi/auth-service/pkg/user_v1"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.runGRPCServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initLogger,
		a.initTracing,
		a.initJwt,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initLogger(_ context.Context) error {
	logger.Init(
		logger.GetCore(
			logger.GetAtomicLevel(a.serviceProvider.LoggerConfig().Level()),
		),
	)
	logger.Info("Logger is running")
	return nil
}

func (a *App) initTracing(_ context.Context) error {
	tracing.Init(a.serviceProvider.ServiceNameConfig().Service())
	logger.Info("Tracing is running")
	return nil
}

func (a *App) initJwt(_ context.Context) error {
	a.serviceProvider.JWTConfig()
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	rateLimiter := rate_limiter.NewTokenBucketLimiter(ctx, 10, time.Second) // Подключаем rate_limiter (ограничение 10 rps в течении секунды)

	cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{ // Создаем Circuit breaker
		Name:        "my-service",    // Имя сервиса, где запускаем Circuit breaker
		MaxRequests: 3,               // Максимально кол-во запросов, которые могут проходить через Circuit breaker когда он в полуоткрытом состоянии
		Timeout:     5 * time.Second, // Период открытой стадии Circuit breaker. Если 5 секунд проходит, то он переходит в полуоткрытое состояние и затем через 5 секунд в закрытое или наоборот
		ReadyToTrip: func(counts gobreaker.Counts) bool { // Ф-ция по какому условию Circuit breaker должен перейти в открытое состояние
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return failureRatio >= 0.6 // Если количество запросов ошибочных относительно общего кол-ва больше или равно 60 %, то открывается Circuit breaker
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) { // Ф-ция когда Circuit breaker будем менять свое состояние
			logger.Warn(fmt.Sprintf("Circuit Breaker: %s, changed from %v, to %v\n", name, from, to))
		},
	})

	a.grpcServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				interceptor.LogInterceptor,
				interceptor.ServerTracingInterceptor,
				interceptor.NewRateLimiterInterceptor(rateLimiter).Unary, // подключаем interceptor rate_limiter
				interceptor.NewCircuitBreakerInterceptor(cb).Unary,       // подключаем interceptor circuit breaker
				interceptor.ErrorCodesInterceptor,
			),
		),
	)

	reflection.Register(a.grpcServer)

	descAuth.RegisterAuthV1Server(a.grpcServer, a.serviceProvider.AuthImpl(ctx))
	descUser.RegisterUserV1Server(a.grpcServer, a.serviceProvider.UserImpl(ctx))

	return nil
}

func (a *App) runGRPCServer() error {
	logger.Info(fmt.Sprintf("GRPC server is running on %s", a.serviceProvider.GRPCConfig().Address()))

	list, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}
