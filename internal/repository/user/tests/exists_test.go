package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jedi/auth/internal/repository"
	repoMocks "github.com/go-jedi/auth/internal/repository/mocks"
	"github.com/golang/mock/gomock"
)

func TestExists(t *testing.T) {
	type userRepositoryMockFunc func(mc *gomock.Controller) repository.UserRepository

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

		repoErr = errors.New("repository error")

		username = gofakeit.Name()
		email    = gofakeit.Email()
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
				ctx:      ctx,
				username: username,
				email:    email,
			},
			want: want{
				exists: true,
				err:    nil,
			},
			userRepositoryMockFunc: func(mc *gomock.Controller) repository.UserRepository {
				mock := repoMocks.NewMockUserRepository(mc)
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
				err:    repoErr,
			},
			userRepositoryMockFunc: func(mc *gomock.Controller) repository.UserRepository {
				mock := repoMocks.NewMockUserRepository(mc)
				mock.EXPECT().Exists(ctx, username, email).Return(false, repoErr)
				return mock
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			urm := test.userRepositoryMockFunc(mc)

			ie, err := urm.Exists(test.in.ctx, test.in.username, test.in.email)

			if !errors.Is(err, test.want.err) {
				t.Fatalf("got %v; want %v", err, test.want.err)
			}
			if ie != test.want.exists {
				t.Fatalf("got %v; want %v", ie, test.want.exists)
			}
		})
	}
}
