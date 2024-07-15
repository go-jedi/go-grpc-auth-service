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

func TestCreate(t *testing.T) {
	type userServiceMockFunc func(mc *gomock.Controller) service.UserService

	mc := gomock.NewController(t)
	defer mc.Finish()

	type in struct {
		ctx context.Context
		dto user.CreateDTO
	}

	type want struct {
		usr user.User
		err error
	}

	var (
		ctx = context.TODO()

		servErr = errors.New("service error")

		id           = int64(1)
		username     = gofakeit.Name()
		fullName     = gofakeit.Name()
		email        = gofakeit.Email()
		password     = gofakeit.Password(true, true, true, true, true, 16)
		passwordHash = gofakeit.Password(true, true, true, true, true, 16)
		createdAt    = time.Now()
		updatedAt    = time.Now()

		dto = user.CreateDTO{
			Username: username,
			FullName: fullName,
			Email:    email,
			Password: password,
		}

		usr = user.User{
			ID:           id,
			Username:     username,
			FullName:     fullName,
			Email:        email,
			Password:     password,
			PasswordHash: passwordHash,
			CreatedAt:    createdAt,
			UpdatedAt:    updatedAt,
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
				dto: dto,
			},
			want: want{
				usr: usr,
				err: nil,
			},
			userServiceMockFunc: func(mc *gomock.Controller) service.UserService {
				mock := servMocks.NewMockUserService(mc)
				mock.EXPECT().Create(ctx, dto).Return(usr, nil)
				return mock
			},
		},
		{
			name: "error",
			in: in{
				ctx: ctx,
				dto: dto,
			},
			want: want{
				usr: usr,
				err: servErr,
			},
			userServiceMockFunc: func(mc *gomock.Controller) service.UserService {
				mock := servMocks.NewMockUserService(mc)
				mock.EXPECT().Create(ctx, dto).Return(usr, servErr)
				return mock
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			usm := test.userServiceMockFunc(mc)

			u, err := usm.Create(test.in.ctx, test.in.dto)

			if !errors.Is(err, test.want.err) {
				t.Fatalf("got %v; want %v", err, test.want.err)
			}
			if !reflect.DeepEqual(u, test.want.usr) {
				t.Fatalf("got user %v; want %v", u, test.want.usr)
			}
		})
	}
}
