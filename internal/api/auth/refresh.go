package auth

import (
	"context"

	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
)

func (i *Implementation) Refresh(context.Context, *protoservice.RefreshRequest) (*protoservice.RefreshResponse, error) {
	return &protoservice.RefreshResponse{}, nil
}
