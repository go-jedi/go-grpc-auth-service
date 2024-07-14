package user

import (
	"context"

	protomodel "github.com/go-jedi/auth/gen/proto/model/v1"
	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
)

func (i *Implementation) Update(context.Context, *protoservice.UpdateRequest) (*protomodel.User, error) {
	return &protomodel.User{}, nil
}
