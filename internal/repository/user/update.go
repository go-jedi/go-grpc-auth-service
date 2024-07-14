package user

import (
	"context"

	"github.com/go-jedi/auth/internal/domain/user"
)

func (r *repo) Update(_ context.Context, _ user.UpdateDTO) (user.User, error) {
	return user.User{}, nil
}
