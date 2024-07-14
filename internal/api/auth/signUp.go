package auth

import (
	"context"

	protomodel "github.com/go-jedi/auth/gen/proto/model/v1"
	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
	"github.com/go-jedi/auth/internal/domain/user"
)

func (i *Implementation) SignUp(ctx context.Context, req *protoservice.SignUpRequest) (*protomodel.User, error) {
	dto := user.CreateDTO{
		Username: req.GetUsername(),
		FullName: req.GetFullName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	// check valid dto
	if err := i.validator.Struct(dto); err != nil {
		return nil, err
	}

	u, err := i.authService.SignUp(ctx, dto)
	if err != nil {
		return nil, err
	}

	return u.ToProto(), nil
}
