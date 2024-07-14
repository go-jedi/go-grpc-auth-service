package user

import (
	"context"

	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
)

func (i *Implementation) Exists(ctx context.Context, in *protoservice.ExistsRequest) (*protoservice.ExistsResponse, error) {
	ie, err := i.userService.Exists(ctx, in.GetUsername(), in.GetEmail())
	if err != nil {
		return nil, err
	}

	return &protoservice.ExistsResponse{
		Exists: ie,
	}, nil
}
