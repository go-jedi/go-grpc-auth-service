package user

import (
	"context"

	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) ExistsUsername(ctx context.Context, in *protoservice.ExistsUsernameRequest) (*protoservice.ExistsUsernameResponse, error) {
	ie, err := h.userService.ExistsUsername(ctx, in.GetUsername())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &protoservice.ExistsUsernameResponse{
		Exists: ie,
	}, nil
}
