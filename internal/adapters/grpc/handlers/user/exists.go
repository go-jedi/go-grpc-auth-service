package user

import (
	"context"

	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) Exists(ctx context.Context, in *protoservice.ExistsRequest) (*protoservice.ExistsResponse, error) {
	ie, err := h.userService.Exists(ctx, in.GetUsername(), in.GetEmail())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &protoservice.ExistsResponse{
		Exists: ie,
	}, nil
}
