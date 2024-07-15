package user

import (
	"context"

	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) ExistsEmail(ctx context.Context, in *protoservice.ExistsEmailRequest) (*protoservice.ExistsEmailResponse, error) {
	ie, err := i.userService.ExistsEmail(ctx, in.GetEmail())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &protoservice.ExistsEmailResponse{
		Exists: ie,
	}, nil
}
