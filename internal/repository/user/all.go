package user

import (
	"context"

	"github.com/go-jedi/auth/internal/domain/user"
)

func (r *repo) All(_ context.Context) ([]user.User, error) {
	return nil, nil
}
