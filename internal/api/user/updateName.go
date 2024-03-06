package user

import (
	"context"

	"github.com/go-jedi/auth-service/internal/utils/jwt"
	"github.com/go-jedi/platform_common/pkg/sys"
	"github.com/go-jedi/platform_common/pkg/sys/codes"

	"github.com/go-jedi/auth-service/internal/converter"

	"github.com/go-jedi/auth-service/internal/logger"
	"go.uber.org/zap"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/go-jedi/auth-service/pkg/user_v1"
)

func (i *Implementation) UpdateName(ctx context.Context, req *desc.UpdateNameRequest) (*emptypb.Empty, error) {
	logger.Info(
		"(API) UpdateName user...",
		zap.Int64("id", req.GetId()),
		zap.String("username", req.GetUsername()),
	)

	resultCheck, err := jwt.Check(ctx)
	if err != nil {
		return nil, err
	}

	if !resultCheck {
		return nil, sys.NewCommonError("authentication error", codes.Unauthenticated)
	}

	err = i.userService.UpdateName(ctx, converter.ToUpdateNameServiceFromProto(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
