package auth

import (
	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
	"github.com/go-jedi/auth/internal/service"
	"github.com/go-jedi/auth/pkg/logger"
	"github.com/go-jedi/auth/pkg/validator"
)

type Implementation struct {
	protoservice.UnimplementedAuthV1Server
	authService service.AuthService
	logger      *logger.Logger
	validator   *validator.Validator
}

func NewImplementation(
	authService service.AuthService,
	logger *logger.Logger,
	validator *validator.Validator,
) *Implementation {
	return &Implementation{
		authService: authService,
		logger:      logger,
		validator:   validator,
	}
}
