package service

import (
	"context"

	"github.com/go-jedi/auth/internal/domain/auth"
	"github.com/go-jedi/auth/internal/domain/user"
)

type AuthService interface {
	SignUp(ctx context.Context, dto user.CreateDTO) (user.User, error)
	SignIn(ctx context.Context, dto auth.SignInDTO) (auth.SignInResp, error)
	Check(ctx context.Context, dto auth.CheckDTO) (auth.CheckResp, error)
	Refresh(ctx context.Context, dto auth.RefreshDTO) (auth.RefreshResp, error)
}

type UserService interface {
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
