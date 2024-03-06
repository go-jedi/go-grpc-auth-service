package repository

import (
	"context"

	"github.com/go-jedi/auth-service/internal/model"
)

type AuthRepository interface {
	Register(ctx context.Context, registerRequest *model.RegisterRequest) error
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
}

type UserRepository interface {
	Get(ctx context.Context, id int64) (*model.User, error)
	UpdateName(ctx context.Context, updateNameRequest *model.UpdateNameRequest) error
	UpdatePassword(ctx context.Context, updatePasswordRequest *model.UpdatePasswordRequest) error
}
