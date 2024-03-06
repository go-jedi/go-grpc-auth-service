package service

import (
	"context"

	"github.com/go-jedi/auth-service/internal/model"
)

type AuthService interface {
	Register(ctx context.Context, registerRequest *model.RegisterRequest) error
	Login(ctx context.Context, loginRequest *model.LoginRequest) (*model.LoginResponse, error)
	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
	GetRefreshToken(ctx context.Context, refreshToken string) (string, error)
}

type UserService interface {
	Get(ctx context.Context, id int64) (*model.User, error)
	UpdateName(ctx context.Context, updateNameRequest *model.UpdateNameRequest) error
	UpdatePassword(ctx context.Context, updatePasswordRequest *model.UpdatePasswordRequest) error
}
