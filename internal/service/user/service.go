package user

import (
	"github.com/go-jedi/auth/internal/repository"
	"github.com/go-jedi/auth/internal/service"
	"github.com/go-jedi/auth/pkg/logger"
)

type serv struct {
	userRepository repository.UserRepository
	logger         *logger.Logger
}

func NewService(
	userRepository repository.UserRepository,
	l *logger.Logger,
) service.UserService {
	return &serv{
		userRepository: userRepository,
		logger:         l,
	}
}
