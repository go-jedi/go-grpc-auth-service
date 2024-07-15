package user

import (
	"github.com/go-jedi/auth/internal/repository"
	"github.com/go-jedi/auth/internal/service"
	"github.com/go-jedi/auth/pkg/logger"
	"github.com/go-jedi/auth/pkg/redis"
)

type serv struct {
	userRepository repository.UserRepository
	logger         *logger.Logger
	cache          *redis.Redis
}

func NewService(
	userRepository repository.UserRepository,
	logger *logger.Logger,
	cache *redis.Redis,
) service.UserService {
	return &serv{
		userRepository: userRepository,
		logger:         logger,
		cache:          cache,
	}
}
