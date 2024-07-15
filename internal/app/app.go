package app

import (
	"context"

	"github.com/go-jedi/auth/config"
	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
	"github.com/go-jedi/auth/pkg/grpcserver"
	"github.com/go-jedi/auth/pkg/jwt"
	"github.com/go-jedi/auth/pkg/logger"
	"github.com/go-jedi/auth/pkg/postgres"
	"github.com/go-jedi/auth/pkg/redis"
	"github.com/go-jedi/auth/pkg/validator"
)

type App struct {
	cfg config.Config

	logger    *logger.Logger
	validator *validator.Validator
	jwt       *jwt.JWT
	gs        *grpcserver.GRPCServer
	db        *postgres.Postgres
	cache     *redis.Redis
	sp        *serviceProvider
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	return a, nil
}

// Run application.
func (a *App) Run() error {
	return a.runGRPCServer()
}

// initDeps initialize deps.
func (a *App) initDeps(ctx context.Context) error {
	i := []func(context.Context) error{
		a.initConfig,
		a.initLogger,
		a.initValidator,
		a.initJWT,
		a.initPostgres,
		a.initRedis,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range i {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

// initConfig initialize config.
func (a *App) initConfig(_ context.Context) (err error) {
	a.cfg, err = config.GetConfig()
	if err != nil {
		return err
	}

	return nil
}

// initLogger initialize logger.
func (a *App) initLogger(_ context.Context) error {
	a.logger = logger.NewLogger(a.cfg.Logger)
	return nil
}

// initValidator initialize validator.
func (a *App) initValidator(_ context.Context) error {
	a.validator = validator.NewValidator()
	return nil
}

// initJWT initialize jwt.
func (a *App) initJWT(_ context.Context) (err error) {
	a.jwt, err = jwt.NewJWT(a.cfg.JWT)
	if err != nil {
		return err
	}

	return nil
}

// initPostgres initialize postgres.
func (a *App) initPostgres(ctx context.Context) (err error) {
	a.db, err = postgres.NewPostgres(ctx, a.cfg.Postgres)
	if err != nil {
		return err
	}

	return nil
}

// initRedis initialize redis.
func (a *App) initRedis(ctx context.Context) (err error) {
	a.cache, err = redis.NewRedis(ctx, a.cfg.Redis)
	if err != nil {
		return err
	}

	return nil
}

// initServiceProvider initialize server provider.
func (a *App) initServiceProvider(_ context.Context) error {
	a.sp = newServiceProvider(a.logger, a.validator, a.jwt, a.db, a.cache)
	return nil
}

// initGRPCServer initialize grpc server.
func (a *App) initGRPCServer(ctx context.Context) (err error) {
	a.gs, err = grpcserver.NewGRPCServer(a.cfg.GRPCServer)
	if err != nil {
		return err
	}

	protoservice.RegisterAuthV1Server(a.gs.Server, a.sp.AuthImpl(ctx))
	protoservice.RegisterUserV1Server(a.gs.Server, a.sp.UserImpl(ctx))

	return nil
}

// runGRPCServer run server.
func (a *App) runGRPCServer() error {
	return a.gs.Start()
}
