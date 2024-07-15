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

func TestExists(t *testing.T) {
	type userServiceMockFunc func(mc *gomock.Controller) service.UserService

	mc := gomock.NewController(t)
	defer mc.Finish()

	type in struct {
		ctx      context.Context
		username string
		email    string
	}

	type want struct {
		exists bool
		err    error
	}

	var (
		ctx = context.TODO()

		servErr = errors.New("service error")

		username = gofakeit.Name()
		email    = gofakeit.Email()
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
				ctx:      ctx,
				username: username,
				email:    email,
			},
			want: want{
				exists: true,
				err:    nil,
			},
			userServiceMockFunc: func(mc *gomock.Controller) service.UserService {
				mock := servMocks.NewMockUserService(mc)
				mock.EXPECT().Exists(ctx, username, email).Return(true, nil)
				return mock
			},
		},
		{
			name: "error",
			in: in{
				ctx:      ctx,
				username: username,
				email:    email,
			},
			want: want{
				exists: false,
				err:    servErr,
			},
			userServiceMockFunc: func(mc *gomock.Controller) service.UserService {
				mock := servMocks.NewMockUserService(mc)
				mock.EXPECT().Exists(ctx, username, email).Return(false, servErr)
				return mock
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			usm := test.userServiceMockFunc(mc)

			ie, err := usm.Exists(test.in.ctx, test.in.username, test.in.email)

			if !errors.Is(err, test.want.err) {
				t.Fatalf("got %v; want %v", err, test.want.err)
			}
			if ie != test.want.exists {
				t.Fatalf("got %v; want %v", ie, test.want.exists)
			}
		})
	}
}
