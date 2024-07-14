package user

import (
	"time"

	protomodel "github.com/go-jedi/auth/gen/proto/model/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type User struct {
	ID           int64
	Username     string
	FullName     string
	Email        string
	Password     string
	PasswordHash string
	Deleted      bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) ToProto() *protomodel.User {
	return &protomodel.User{
		Id:           u.ID,
		Username:     u.Username,
		FullName:     u.FullName,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		Deleted:      u.Deleted,
		CreatedAt:    timestamppb.New(u.CreatedAt),
		UpdatedAt:    timestamppb.New(u.UpdatedAt),
	}
}

func SliceToProto(in []User) []*protomodel.User {
	out := make([]*protomodel.User, 0, len(in))

	for _, u := range in {
		out = append(out, u.ToProto())
	}

	return out
}

//
// CREATE
//

type CreateDTO struct {
	Username string `validate:"required"`
	FullName string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required,min=8,max=32"`
}

//
// UPDATE
//

type UpdateDTO struct {
	ID       int64  `validate:"required,min=1"`
	Username string `validate:"required"`
	FullName string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required,min=8"`
}
