package user

import (
	"context"

	protomodel "github.com/go-jedi/auth/gen/proto/model/v1"
	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
)

func (i *Implementation) GetByID(ctx context.Context, in *protoservice.GetByIDRequest) (*protomodel.User, error) {
	u, err := i.userService.GetByID(ctx, in.GetId())
	if err != nil {
		return nil, err
	}

	return u.ToProto(), nil
}
