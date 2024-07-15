package user

import (
	"context"

	protomodel "github.com/go-jedi/auth/gen/proto/model/v1"
	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
	"github.com/go-jedi/auth/internal/domain/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Update(ctx context.Context, in *protoservice.UpdateRequest) (*protomodel.User, error) {
	dto := user.UpdateDTO{
		ID:       in.GetId(),
		Username: in.GetUsername(),
		FullName: in.GetFullName(),
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
	}

	// check valid dto
	if err := i.validator.Struct(dto); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	u, err := i.userService.Update(ctx, dto)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return u.ToProto(), nil
}
