package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/go-jedi/auth/internal/service"
	servMocks "github.com/go-jedi/auth/internal/service/mocks"
	"github.com/golang/mock/gomock"
)

func TestDelete(t *testing.T) {
	type userServiceMockFunc func(mc *gomock.Controller) service.UserService

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

		servErr = errors.New("service error")

		id = int64(1)
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
				err: nil,
			},
			userServiceMockFunc: func(mc *gomock.Controller) service.UserService {
				mock := servMocks.NewMockUserService(mc)
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
				err: servErr,
			},
			userServiceMockFunc: func(mc *gomock.Controller) service.UserService {
				mock := servMocks.NewMockUserService(mc)
				mock.EXPECT().Delete(ctx, id).Return(servErr)
				return mock
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			usm := test.userServiceMockFunc(mc)

			err := usm.Delete(test.in.ctx, test.in.id)

			if !errors.Is(err, test.want.err) {
				t.Fatalf("got %v; want %v", err, test.want.err)
			}
		})
	}
}
