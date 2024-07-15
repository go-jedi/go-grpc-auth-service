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

func TestAll(t *testing.T) {
	type userRepositoryMockFunc func(mc *gomock.Controller) repository.UserRepository

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

		repoErr = errors.New("repository error")

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
		name                   string
		in                     in
		want                   want
		userRepositoryMockFunc userRepositoryMockFunc
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
			userRepositoryMockFunc: func(mc *gomock.Controller) repository.UserRepository {
				mock := repoMocks.NewMockUserRepository(mc)
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
				users: nil,
				err:   repoErr,
			},
			userRepositoryMockFunc: func(mc *gomock.Controller) repository.UserRepository {
				mock := repoMocks.NewMockUserRepository(mc)
				mock.EXPECT().All(ctx).Return(nil, repoErr)
				return mock
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			urm := test.userRepositoryMockFunc(mc)

			u, err := urm.All(test.in.ctx)

			if !errors.Is(err, test.want.err) {
				t.Fatalf("got %v; want %v", err, test.want.err)
			}
			if !reflect.DeepEqual(u, test.want.users) {
				t.Fatalf("got users %v; want %v", u, test.want.users)
			}
		})
	}
}
