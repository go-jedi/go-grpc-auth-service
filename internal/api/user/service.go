package user

import (
	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
	"github.com/go-jedi/auth/internal/service"
	"github.com/go-jedi/auth/pkg/logger"
	"github.com/go-jedi/auth/pkg/validator"
)

type Implementation struct {
	protoservice.UnimplementedUserV1Server
	userService service.UserService
	logger      *logger.Logger
	validator   *validator.Validator
}

func NewImplementation(
	userService service.UserService,
	logger *logger.Logger,
	validator *validator.Validator,
) *Implementation {
	return &Implementation{
		userService: userService,
		logger:      logger,
		validator:   validator,
	}
}
