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

func (i *Implementation) UpdatePassword(ctx context.Context, req *desc.UpdatePasswordRequest) (*emptypb.Empty, error) {
	logger.Info(
		"(API) UpdatePassword user...",
		zap.Int64("id", req.GetId()),
		zap.String("password", req.GetPassword()),
	)

	resultCheck, err := jwt.Check(ctx)
	if err != nil {
		return nil, err
	}

	if !resultCheck {
		return nil, sys.NewCommonError("authentication error", codes.Unauthenticated)
	}

	err = i.userService.UpdatePassword(ctx, converter.ToUpdatePasswordServiceFromProto(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
