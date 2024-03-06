package auth

import (
	"context"

	"go.uber.org/zap"

	"github.com/go-jedi/auth-service/internal/converter"
	"github.com/go-jedi/auth-service/internal/logger"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/go-jedi/auth-service/pkg/auth_v1"
)

func (i *Implementation) Register(ctx context.Context, req *desc.RegisterRequest) (*emptypb.Empty, error) {
	logger.Info(
		"(API) Register auth...",
		zap.String("username", req.GetUsername()),
		zap.String("password", req.GetPassword()),
	)

	err := i.authService.Register(ctx, converter.ToRegisterServiceFromProto(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
