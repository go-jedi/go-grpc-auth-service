package auth

import (
	"context"

	"github.com/go-jedi/auth-service/internal/logger"
	"go.uber.org/zap"

	desc "github.com/go-jedi/auth-service/pkg/auth_v1"
)

func (i *Implementation) GetAccessToken(ctx context.Context, req *desc.GetAccessTokenRequest) (*desc.GetAccessTokenResponse, error) {
	logger.Info(
		"(API) GetAccessToken auth...",
		zap.String("refresh_token", req.RefreshToken),
	)

	accessToken, err := i.authService.GetAccessToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, err
	}

	return &desc.GetAccessTokenResponse{
		AccessToken: accessToken,
	}, nil
}
