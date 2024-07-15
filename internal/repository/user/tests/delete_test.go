package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/go-jedi/auth/internal/repository"
	repoMocks "github.com/go-jedi/auth/internal/repository/mocks"
	"github.com/golang/mock/gomock"
)

func TestDelete(t *testing.T) {
	type userRepositoryMockFunc func(mc *gomock.Controller) repository.UserRepository

	mc := gomock.NewController(t)
	defer mc.Finish()

	type in struct {
		ctx context.Context
		id  int64
	}

	type want struct {
		err error
	}

	var (
		ctx = context.TODO()

		repoErr = errors.New("repository error")

		id = int64(1)
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
				err: nil,
			},
			userRepositoryMockFunc: func(mc *gomock.Controller) repository.UserRepository {
				mock := repoMocks.NewMockUserRepository(mc)
				mock.EXPECT().Delete(ctx, id).Return(nil)
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
				err: repoErr,
			},
			userRepositoryMockFunc: func(mc *gomock.Controller) repository.UserRepository {
				mock := repoMocks.NewMockUserRepository(mc)
				mock.EXPECT().Delete(ctx, id).Return(repoErr)
				return mock
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			urm := test.userRepositoryMockFunc(mc)

			err := urm.Delete(test.in.ctx, test.in.id)

			if !errors.Is(err, test.want.err) {
				t.Fatalf("got %v; want %v", err, test.want.err)
			}
		})
	}
}
