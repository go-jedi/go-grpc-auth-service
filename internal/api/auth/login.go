package auth

import (
	"context"

	"github.com/go-jedi/auth-service/internal/converter"

	"github.com/go-jedi/auth-service/internal/logger"
	desc "github.com/go-jedi/auth-service/pkg/auth_v1"
	"go.uber.org/zap"
)

func (i *Implementation) Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {
	logger.Info(
		"(API) Login auth...",
		zap.String("username", req.Username),
		zap.String("password", req.Password),
	)

	result, err := i.authService.Login(ctx, converter.ToLoginServiceFromProto(req))
	if err != nil {
		return nil, err
	}

	return converter.ToLoginProtoFromService(result), nil
}
