package auth

import (
	"context"

	"github.com/go-jedi/auth-service/internal/logger"
	"go.uber.org/zap"

	desc "github.com/go-jedi/auth-service/pkg/auth_v1"
)

func (i *Implementation) GetRefreshToken(ctx context.Context, req *desc.GetRefreshTokenRequest) (*desc.GetRefreshTokenResponse, error) {
	logger.Info(
		"(API) GetRefreshToken auth...",
		zap.String("refresh_token", req.RefreshToken),
	)

	refreshToken, err := i.authService.GetRefreshToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, err
	}

	return &desc.GetRefreshTokenResponse{
		RefreshToken: refreshToken,
	}, nil
}
