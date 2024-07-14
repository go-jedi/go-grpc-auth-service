package app

import (
	"context"

	"github.com/go-jedi/auth/internal/api/auth"
	"github.com/go-jedi/auth/internal/api/user"
	"github.com/go-jedi/auth/internal/repository"
	userRepository "github.com/go-jedi/auth/internal/repository/user"
	"github.com/go-jedi/auth/internal/service"
	authService "github.com/go-jedi/auth/internal/service/auth"
	userService "github.com/go-jedi/auth/internal/service/user"
	"github.com/go-jedi/auth/pkg/jwt"
	"github.com/go-jedi/auth/pkg/logger"
	"github.com/go-jedi/auth/pkg/postgres"
	"github.com/go-jedi/auth/pkg/validator"
)

type serviceProvider struct {
	logger    *logger.Logger
	validator *validator.Validator
	jwt       *jwt.JWT
	db        *postgres.Postgres

	// auth
	authService service.AuthService
	authImpl    *auth.Implementation

	// user
	userRepository repository.UserRepository
	userService    service.UserService
	userImpl       *user.Implementation
}

func newServiceProvider(
	logger *logger.Logger,
	validator *validator.Validator,
	jwt *jwt.JWT,
	db *postgres.Postgres,
) *serviceProvider {
	return &serviceProvider{
		logger:    logger,
		validator: validator,
		jwt:       jwt,
		db:        db,
	}
}

//
// AUTH
//

func (sp *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if sp.authService == nil {
		sp.authService = authService.NewService(
			sp.UserRepository(ctx),
			sp.logger,
			sp.jwt,
		)
	}

	return sp.authService
}

func (sp *serviceProvider) AuthImpl(ctx context.Context) *auth.Implementation {
	if sp.authImpl == nil {
		sp.authImpl = auth.NewImplementation(
			sp.AuthService(ctx),
			sp.logger,
			sp.validator,
		)
	}

	return sp.authImpl
}

//
// USER
//

func (sp *serviceProvider) UserRepository(_ context.Context) repository.UserRepository {
	if sp.userRepository == nil {
		sp.userRepository = userRepository.NewRepository(
			sp.logger,
			sp.db,
		)
	}

	return sp.userRepository
}

func (sp *serviceProvider) UserService(ctx context.Context) service.UserService {
	if sp.userService == nil {
		sp.userService = userService.NewService(
			sp.UserRepository(ctx),
			sp.logger,
		)
	}

	return sp.userService
}

func (sp *serviceProvider) UserImpl(ctx context.Context) *user.Implementation {
	if sp.userImpl == nil {
		sp.userImpl = user.NewImplementation(
			sp.UserService(ctx),
			sp.logger,
			sp.validator,
		)
	}

	return sp.userImpl
}
