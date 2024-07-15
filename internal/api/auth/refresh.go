package auth

import (
	"context"

	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
	"github.com/go-jedi/auth/internal/domain/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Refresh(ctx context.Context, in *protoservice.RefreshRequest) (*protoservice.RefreshResponse, error) {
	dto := auth.RefreshDTO{
		ID:           in.GetId(),
		RefreshToken: in.GetRefreshToken(),
	}

	// check valid dto
	if err := i.validator.Struct(dto); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	r, err := i.authService.Refresh(ctx, dto)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return r.ToProto(), nil
}
