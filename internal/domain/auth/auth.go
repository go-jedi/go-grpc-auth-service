package auth

import (
	"time"

	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//
// SIGN IN
//

type SignInDTO struct {
	Username string `validate:"required"`
	Password string `validate:"required,min=8,max=32"`
}

type SignInResp struct {
	AccessToken  string
	RefreshToken string
	AccessExpAt  time.Time
	RefreshExpAt time.Time
}

func (sir *SignInResp) ToProto() *protoservice.SignInResponse {
	return &protoservice.SignInResponse{
		AccessToken:  sir.AccessToken,
		RefreshToken: sir.RefreshToken,
		AccessExpAt:  timestamppb.New(sir.AccessExpAt),
		RefreshExpAt: timestamppb.New(sir.RefreshExpAt),
	}
}

//
// CHECK
//

type CheckDTO struct {
	ID    int64  `validate:"required,min=1"`
	Token string `validate:"required"`
}

type CheckResp struct {
	ID       int64
	Username string
	Token    string
	ExpAt    time.Time
}

func (cr *CheckResp) ToProto() *protoservice.CheckResponse {
	return &protoservice.CheckResponse{
		Id:       cr.ID,
		Username: cr.Username,
		Token:    cr.Token,
		ExpAt:    timestamppb.New(cr.ExpAt),
	}
}

//
// REFRESH
//

type RefreshDTO struct {
	ID           int64  `validate:"required,min=1"`
	RefreshToken string `validate:"required"`
}

type RefreshResp struct {
	AccessToken  string
	RefreshToken string
	AccessExpAt  time.Time
	RefreshExpAt time.Time
}

func (rr *RefreshResp) ToProto() *protoservice.RefreshResponse {
	return &protoservice.RefreshResponse{
		AccessToken:  rr.AccessToken,
		RefreshToken: rr.RefreshToken,
		AccessExpAt:  timestamppb.New(rr.AccessExpAt),
		RefreshExpAt: timestamppb.New(rr.RefreshExpAt),
	}
}
