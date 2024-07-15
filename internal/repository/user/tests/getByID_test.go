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

func TestGetByID(t *testing.T) {
	type userRepositoryMockFunc func(mc *gomock.Controller) repository.UserRepository

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

		repoErr = errors.New("repository error")

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
		name                   string
		in                     in
		want                   want
		userRepositoryMockFunc userRepositoryMockFunc
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
			userRepositoryMockFunc: func(mc *gomock.Controller) repository.UserRepository {
				mock := repoMocks.NewMockUserRepository(mc)
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
				err: repoErr,
			},
			userRepositoryMockFunc: func(mc *gomock.Controller) repository.UserRepository {
				mock := repoMocks.NewMockUserRepository(mc)
				mock.EXPECT().GetByID(ctx, id).Return(usr, repoErr)
				return mock
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			urm := test.userRepositoryMockFunc(mc)

			u, err := urm.GetByID(test.in.ctx, test.in.id)

			if !errors.Is(err, test.want.err) {
				t.Fatalf("got %v; want %v", err, test.want.err)
			}
			if !reflect.DeepEqual(u, test.want.usr) {
				t.Fatalf("got user %v; want %v", u, test.want.usr)
			}
		})
	}
}
