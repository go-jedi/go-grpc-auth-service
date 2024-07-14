package user

import (
	"context"

	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
)

func (i *Implementation) Exists(context.Context, *protoservice.ExistsRequest) (*protoservice.ExistsResponse, error) {
	return &protoservice.ExistsResponse{}, nil
}
