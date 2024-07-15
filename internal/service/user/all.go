package user

import (
	"context"
	"log"

	"github.com/go-jedi/auth/internal/domain/user"
)

func (s *serv) All(ctx context.Context) ([]user.User, error) {
	return s.allFromCacheOrRepo(ctx)
}

// allFromCacheOrRepo get all users from cache or repository.
func (s *serv) allFromCacheOrRepo(ctx context.Context) ([]user.User, error) {
	u, err := s.cache.User.All(ctx)
	if err != nil {
		log.Println(err)
		return s.userRepository.All(ctx)
	}

	if len(u) == 0 {
		return s.userRepository.All(ctx)
	}

	return u, nil
}
