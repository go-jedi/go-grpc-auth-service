package user

import (
	"context"

	protomodel "github.com/go-jedi/auth/gen/proto/model/v1"
	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
	"github.com/go-jedi/auth/internal/domain/user"
)

func (i *Implementation) Update(ctx context.Context, req *protoservice.UpdateRequest) (*protomodel.User, error) {
	dto := user.UpdateDTO{
		ID:       req.GetId(),
		Username: req.GetUsername(),
		FullName: req.GetFullName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	u, err := i.userService.Update(ctx, dto)
	if err != nil {
		return nil, err
	}

	return u.ToProto(), nil
}
