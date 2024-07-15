package tests

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jedi/auth/internal/domain/user"
	"github.com/go-jedi/auth/internal/service"
	servMocks "github.com/go-jedi/auth/internal/service/mocks"
	"github.com/golang/mock/gomock"
)

func TestAll(t *testing.T) {
	type userServiceMockFunc func(mc *gomock.Controller) service.UserService

	mc := gomock.NewController(t)
	defer mc.Finish()

	type in struct {
		ctx context.Context
	}

	type want struct {
		users []user.User
		err   error
	}

	var (
		ctx = context.TODO()

		servErr = errors.New("service error")

		usrs = []user.User{
			{
				ID:           1,
				Username:     gofakeit.Name(),
				FullName:     gofakeit.Name(),
				Email:        gofakeit.Email(),
				PasswordHash: gofakeit.Password(true, true, true, true, true, 16),
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
		}
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
				ctx: ctx,
			},
			want: want{
				users: usrs,
				err:   nil,
			},
			userServiceMockFunc: func(mc *gomock.Controller) service.UserService {
				mock := servMocks.NewMockUserService(mc)
				mock.EXPECT().All(ctx).Return(usrs, nil)
				return mock
			},
		},
		{
			name: "error",
			in: in{
				ctx: ctx,
			},
			want: want{
				users: usrs,
				err:   servErr,
			},
			userServiceMockFunc: func(mc *gomock.Controller) service.UserService {
				mock := servMocks.NewMockUserService(mc)
				mock.EXPECT().All(ctx).Return(usrs, servErr)
				return mock
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			usm := test.userServiceMockFunc(mc)

			u, err := usm.All(test.in.ctx)

			if !errors.Is(err, test.want.err) {
				t.Fatalf("got %v; want %v", err, test.want.err)
			}
			if !reflect.DeepEqual(u, test.want.users) {
				t.Fatalf("got users %v; want %v", u, test.want.users)
			}
		})
	}
}
