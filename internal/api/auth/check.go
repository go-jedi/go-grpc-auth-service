package auth

import (
	"context"

	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
	"github.com/go-jedi/auth/internal/domain/auth"
)

func (i *Implementation) Check(ctx context.Context, in *protoservice.CheckRequest) (*protoservice.CheckResponse, error) {
	dto := auth.CheckDTO{
		ID:    in.GetId(),
		Token: in.GetToken(),
	}

	// check valid dto
	if err := i.validator.Struct(dto); err != nil {
		return nil, err
	}

	r, err := i.authService.Check(ctx, dto)
	if err != nil {
		return nil, err
	}

	return r.ToProto(), nil
}
