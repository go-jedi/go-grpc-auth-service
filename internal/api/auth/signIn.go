package auth

import (
	"context"

	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
	"github.com/go-jedi/auth/internal/domain/auth"
)

func (i *Implementation) SignIn(ctx context.Context, req *protoservice.SignInRequest) (*protoservice.SignInResponse, error) {
	dto := auth.SignInDTO{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	}

	// check valid dto
	if err := i.validator.Struct(dto); err != nil {
		return nil, err
	}

	r, err := i.authService.SignIn(ctx, dto)
	if err != nil {
		return nil, err
	}

	return r.ToProto(), nil
}
