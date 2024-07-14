package auth

import (
	"context"

	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
	"github.com/go-jedi/auth/internal/domain/auth"
)

func (i *Implementation) Refresh(ctx context.Context, req *protoservice.RefreshRequest) (*protoservice.RefreshResponse, error) {
	dto := auth.RefreshDTO{
		ID:           req.GetId(),
		RefreshToken: req.GetRefreshToken(),
	}

	// check valid dto
	if err := i.validator.Struct(dto); err != nil {
		return nil, err
	}

	r, err := i.authService.Refresh(ctx, dto)
	if err != nil {
		return nil, err
	}

	return r.ToProto(), nil
}
