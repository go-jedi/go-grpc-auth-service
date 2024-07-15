package user

import (
	"context"
	"log"
	"strconv"

	"github.com/go-jedi/auth/internal/domain/user"
)

func (s *serv) GetByID(ctx context.Context, id int64) (user.User, error) {
	return s.getByIDFromCacheOrRepo(ctx, id)
}

// getByIDFromCacheOrRepo get user from cache or repository.
func (s *serv) getByIDFromCacheOrRepo(ctx context.Context, id int64) (user.User, error) {
	key := strconv.FormatInt(id, 10)

	u, err := s.cache.User.Get(ctx, key)
	if err != nil {
		log.Println(err)
		return s.userRepository.GetByID(ctx, id)
	}

	return u, nil
}
