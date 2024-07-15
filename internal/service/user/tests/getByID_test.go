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

func TestGetByID(t *testing.T) {
	type userServiceMockFunc func(mc *gomock.Controller) service.UserService

	mc := gomock.NewController(t)
	defer mc.Finish()

	type in struct {
		ctx context.Context
		id  int64
	}

	type want struct {
		usr user.User
		err error
	}

	var (
		ctx = context.TODO()

		servErr = errors.New("service error")

		id = int64(1)

		usr = user.User{
			ID:           id,
			Username:     gofakeit.Name(),
			FullName:     gofakeit.Name(),
			Email:        gofakeit.Email(),
			Password:     gofakeit.Password(true, true, true, true, true, 16),
			PasswordHash: gofakeit.Password(true, true, true, true, true, 16),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
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
				id:  id,
			},
			want: want{
				usr: usr,
				err: nil,
			},
			userServiceMockFunc: func(mc *gomock.Controller) service.UserService {
				mock := servMocks.NewMockUserService(mc)
				mock.EXPECT().GetByID(ctx, id).Return(usr, nil)
				return mock
			},
		},
		{
			name: "error",
			in: in{
				ctx: ctx,
				id:  id,
			},
			want: want{
				usr: usr,
				err: servErr,
			},
			userServiceMockFunc: func(mc *gomock.Controller) service.UserService {
				mock := servMocks.NewMockUserService(mc)
				mock.EXPECT().GetByID(ctx, id).Return(usr, servErr)
				return mock
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			usm := test.userServiceMockFunc(mc)

			u, err := usm.GetByID(test.in.ctx, test.in.id)

			if !errors.Is(err, test.want.err) {
				t.Fatalf("got %v; want %v", err, test.want.err)
			}
			if !reflect.DeepEqual(u, test.want.usr) {
				t.Fatalf("got user %v; want %v", u, test.want.usr)
			}
		})
	}
}
