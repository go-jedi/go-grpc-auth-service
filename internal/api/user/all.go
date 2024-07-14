package user

import (
	"context"

	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
	"github.com/go-jedi/auth/internal/domain/user"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) All(ctx context.Context, _ *emptypb.Empty) (*protoservice.AllResponse, error) {
	u, err := i.userService.All(ctx)
	if err != nil {
		return nil, err
	}

	return &protoservice.AllResponse{
		Users: user.SliceToProto(u),
	}, nil
}
