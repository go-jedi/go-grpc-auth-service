package user

import (
	"context"

	"github.com/go-jedi/auth/internal/domain/user"
)

func (s *serv) GetByUsername(ctx context.Context, username string) (user.User, error) {
	return s.userRepository.GetByUsername(ctx, username)
}
