package user

import (
	"context"

	protomodel "github.com/go-jedi/auth/gen/proto/model/v1"
	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
	"github.com/go-jedi/auth/internal/domain/user"
)

func (i *Implementation) Create(ctx context.Context, in *protoservice.CreateRequest) (*protomodel.User, error) {
	dto := user.CreateDTO{
		Username: in.GetUsername(),
		FullName: in.GetFullName(),
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
	}

	// check valid dto
	if err := i.validator.Struct(dto); err != nil {
		return nil, err
	}

	u, err := i.userService.Create(ctx, dto)
	if err != nil {
		return nil, err
	}

	return u.ToProto(), nil
}
