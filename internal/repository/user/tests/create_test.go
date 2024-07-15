package tests

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jedi/auth/internal/domain/user"
	"github.com/go-jedi/auth/internal/repository"
	repoMocks "github.com/go-jedi/auth/internal/repository/mocks"
	"github.com/golang/mock/gomock"
)

func TestCreate(t *testing.T) {
	type userRepositoryMockFunc func(mc *gomock.Controller) repository.UserRepository

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

		repoErr = errors.New("repository error")

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
		name                   string
		in                     in
		want                   want
		userRepositoryMockFunc userRepositoryMockFunc
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
			userRepositoryMockFunc: func(mc *gomock.Controller) repository.UserRepository {
				mock := repoMocks.NewMockUserRepository(mc)
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
				err: repoErr,
			},
			userRepositoryMockFunc: func(mc *gomock.Controller) repository.UserRepository {
				mock := repoMocks.NewMockUserRepository(mc)
				mock.EXPECT().Create(ctx, dto).Return(usr, repoErr)
				return mock
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			urm := test.userRepositoryMockFunc(mc)

			u, err := urm.Create(test.in.ctx, test.in.dto)

			if !errors.Is(err, test.want.err) {
				t.Fatalf("got %v; want %v", err, test.want.err)
			}
			if !reflect.DeepEqual(u, test.want.usr) {
				t.Fatalf("got user %v; want %v", u, test.want.usr)
			}
		})
	}
}
