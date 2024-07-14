package user

import (
	"context"

	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) All(context.Context, *emptypb.Empty) (*protoservice.AllResponse, error) {
	return &protoservice.AllResponse{}, nil
}
