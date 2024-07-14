package auth

import (
	"github.com/go-jedi/auth/internal/repository"
	"github.com/go-jedi/auth/internal/service"
	"github.com/go-jedi/auth/pkg/jwt"
	"github.com/go-jedi/auth/pkg/logger"
)

type serv struct {
	userRepository repository.UserRepository
	logger         *logger.Logger
	jwt            *jwt.JWT
}

func NewService(
	userRepository repository.UserRepository,
	l *logger.Logger,
	jwt *jwt.JWT,
) service.AuthService {
	return &serv{
		userRepository: userRepository,
		logger:         l,
		jwt:            jwt,
	}
}
