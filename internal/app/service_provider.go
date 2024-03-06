package app

import (
	"context"
	"log"

	"github.com/go-jedi/auth-service/internal/api/user"

	"github.com/go-jedi/platform_common/pkg/closer"
	"github.com/go-jedi/platform_common/pkg/db"
	"github.com/go-jedi/platform_common/pkg/db/pg"
	"github.com/go-jedi/platform_common/pkg/db/transaction"

	"github.com/go-jedi/auth-service/internal/api/auth"
	"github.com/go-jedi/auth-service/internal/config"
	"github.com/go-jedi/auth-service/internal/repository"
	"github.com/go-jedi/auth-service/internal/service"

	authRepository "github.com/go-jedi/auth-service/internal/repository/auth"
	userRepository "github.com/go-jedi/auth-service/internal/repository/user"
	authService "github.com/go-jedi/auth-service/internal/service/auth"
	userService "github.com/go-jedi/auth-service/internal/service/user"
)

type serviceProvider struct {
	pgConfig          config.PGConfig
	loggerConfig      config.LoggerConfig
	serviceNameConfig config.ServiceNameConfig
	grpcConfig        config.GRPCConfig
	jwtConfig         config.JWTConfig

	dbClient  db.Client
	txManager db.TxManager

	authRepository repository.AuthRepository
	userRepository repository.UserRepository

	authService service.AuthService
	userService service.UserService

	authImpl *auth.Implementation
	userImpl *user.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) LoggerConfig() config.LoggerConfig {
	if s.loggerConfig == nil {
		cfg, err := config.NewLoggerConfig()
		if err != nil {
			log.Fatalf("failed to get logger config: %s", err.Error())
		}

		s.loggerConfig = cfg
	}

	return s.loggerConfig
}

func (s *serviceProvider) ServiceNameConfig() config.ServiceNameConfig {
	if s.serviceNameConfig == nil {
		cfg, err := config.NewServiceNameConfig()
		if err != nil {
			log.Fatalf("failed to get service name config: %s", err.Error())
		}
		s.serviceNameConfig = cfg
	}

	return s.serviceNameConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) JWTConfig() config.JWTConfig {
	if s.jwtConfig == nil {
		cfg, err := config.NewJWTConfig()
		if err != nil {
			log.Fatalf("failed to get jwt config: %s", err.Error())
		}

		s.jwtConfig = cfg
	}

	return s.jwtConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) AuthRepository(ctx context.Context) repository.AuthRepository {
	if s.authRepository == nil {
		s.authRepository = authRepository.NewRepository(s.DBClient(ctx))
	}

	return s.authRepository
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewService(
			s.AuthRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.authService
}

func (s *serviceProvider) AuthImpl(ctx context.Context) *auth.Implementation {
	if s.authImpl == nil {
		s.authImpl = auth.NewImplementation(s.AuthService(ctx))
	}

	return s.authImpl
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.DBClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.userService
}

func (s *serviceProvider) UserImpl(ctx context.Context) *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(s.UserService(ctx))
	}

	return s.userImpl
}
