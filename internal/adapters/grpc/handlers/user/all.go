package user

import (
	"context"

	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
	"github.com/go-jedi/auth/internal/domain/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Handler) All(ctx context.Context, _ *emptypb.Empty) (*protoservice.AllResponse, error) {
	u, err := h.userService.All(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &protoservice.AllResponse{
		Users: user.SliceToProto(u),
	}, nil
}
