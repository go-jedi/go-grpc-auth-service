package user

import (
	"context"

	"github.com/go-jedi/auth-service/internal/converter"
	"github.com/go-jedi/auth-service/internal/utils/jwt"
	"github.com/go-jedi/platform_common/pkg/sys"
	"github.com/go-jedi/platform_common/pkg/sys/codes"

	"github.com/go-jedi/auth-service/internal/logger"
	"go.uber.org/zap"

	desc "github.com/go-jedi/auth-service/pkg/user_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	logger.Info(
		"(API) Get user...",
		zap.Int64("id", req.GetId()),
	)

	resultCheck, err := jwt.Check(ctx)
	if err != nil {
		return nil, err
	}

	if !resultCheck {
		return nil, sys.NewCommonError("authentication error", codes.Unauthenticated)
	}

	result, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return converter.ToGetProtoFromService(result), nil
}
