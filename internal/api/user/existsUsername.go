package user

import (
	"context"

	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
)

func (i *Implementation) ExistsUsername(ctx context.Context, in *protoservice.ExistsUsernameRequest) (*protoservice.ExistsUsernameResponse, error) {
	ie, err := i.userService.ExistsUsername(ctx, in.GetUsername())
	if err != nil {
		return nil, err
	}

	return &protoservice.ExistsUsernameResponse{
		Exists: ie,
	}, nil
}
