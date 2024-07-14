package user

import (
	"context"

	"github.com/go-jedi/auth/internal/domain/user"
)

func (s *serv) All(ctx context.Context) ([]user.User, error) {
	return s.userRepository.All(ctx)
}
