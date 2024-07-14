package auth

import (
	"context"

	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
)

func (i *Implementation) Check(context.Context, *protoservice.CheckRequest) (*protoservice.CheckResponse, error) {
	return &protoservice.CheckResponse{}, nil
}
