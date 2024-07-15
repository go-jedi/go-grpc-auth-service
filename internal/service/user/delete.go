package user

import (
	"context"
	"strconv"
)

func (s *serv) Delete(ctx context.Context, id int64) error {
	if err := s.userRepository.Delete(ctx, id); err != nil {
		return err
	}

	// delete user from cache
	if err := s.deleteFromCache(ctx, id); err != nil {
		return err
	}

	return nil
}

// deleteFromCache delete val from cache.
func (s *serv) deleteFromCache(ctx context.Context, id int64) error {
	key := strconv.FormatInt(id, 10)

	return s.cache.User.Del(ctx, key)
}
