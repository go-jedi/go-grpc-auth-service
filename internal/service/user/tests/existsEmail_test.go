package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jedi/auth/internal/service"
	servMocks "github.com/go-jedi/auth/internal/service/mocks"
	"github.com/golang/mock/gomock"
)

func TestExistsEmail(t *testing.T) {
	type userServiceMockFunc func(mc *gomock.Controller) service.UserService

	mc := gomock.NewController(t)
	defer mc.Finish()

	type in struct {
		ctx   context.Context
		email string
	}

	type want struct {
		exists bool
		err    error
	}

	var (
		ctx = context.TODO()

		servErr = errors.New("service error")

		email = gofakeit.Email()
	)

	tests := []struct {
		name                string
		in                  in
		want                want
		userServiceMockFunc userServiceMockFunc
	}{
		{
			name: "ok",
			in: in{
				ctx:   ctx,
				email: email,
			},
			want: want{
				exists: false,
				err:    nil,
			},
			userServiceMockFunc: func(mc *gomock.Controller) service.UserService {
				mock := servMocks.NewMockUserService(mc)
				mock.EXPECT().ExistsEmail(ctx, email).Return(false, nil)
				return mock
			},
		},
		{
			name: "error",
			in: in{
				ctx:   ctx,
				email: email,
			},
			want: want{
				exists: false,
				err:    servErr,
			},
			userServiceMockFunc: func(mc *gomock.Controller) service.UserService {
				mock := servMocks.NewMockUserService(mc)
				mock.EXPECT().ExistsEmail(ctx, email).Return(false, servErr)
				return mock
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			usm := test.userServiceMockFunc(mc)

			ie, err := usm.ExistsEmail(test.in.ctx, test.in.email)

			if !errors.Is(err, test.want.err) {
				t.Fatalf("got %v; want %v", err, test.want.err)
			}
			if ie != test.want.exists {
				t.Fatalf("got %v; want %v", ie, test.want.exists)
			}
		})
	}
}
