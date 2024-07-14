package user

import (
	"context"

	protomodel "github.com/go-jedi/auth/gen/proto/model/v1"
	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
)

func (i *Implementation) GetByUsername(ctx context.Context, req *protoservice.GetByUsernameRequest) (*protomodel.User, error) {
	u, err := i.userService.GetByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, err
	}

	return u.ToProto(), nil
}
