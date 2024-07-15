package user

import (
	"context"

	protomodel "github.com/go-jedi/auth/gen/proto/model/v1"
	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetByUsername(ctx context.Context, in *protoservice.GetByUsernameRequest) (*protomodel.User, error) {
	u, err := i.userService.GetByUsername(ctx, in.GetUsername())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return u.ToProto(), nil
}
