package repository

import (
	"context"

	"github.com/go-jedi/auth/internal/domain/user"
)

type UserRepository interface {
	Create(ctx context.Context, dto user.CreateDTO) (user.User, error)
	All(ctx context.Context) ([]user.User, error)
	GetByID(ctx context.Context, id int64) (user.User, error)
	GetByUsername(ctx context.Context, username string) (user.User, error)
	Exists(ctx context.Context, username string, email string) (bool, error)
	ExistsUsername(ctx context.Context, username string) (bool, error)
	ExistsEmail(ctx context.Context, email string) (bool, error)
	Update(ctx context.Context, dto user.UpdateDTO) (user.User, error)
	Delete(ctx context.Context, id int64) error
}
