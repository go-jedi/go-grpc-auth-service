package user

import (
	"context"

	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(context.Context, *protoservice.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
