package user

import (
	"context"

	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, in *protoservice.DeleteRequest) (*emptypb.Empty, error) {
	if err := i.userService.Delete(ctx, in.GetId()); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
