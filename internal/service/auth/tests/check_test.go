package tests

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jedi/auth/internal/domain/auth"
	"github.com/go-jedi/auth/internal/service"
	servMocks "github.com/go-jedi/auth/internal/service/mocks"
	"github.com/golang/mock/gomock"
)

func TestCheck(t *testing.T) {
	type authServiceMockFunc func(mc *gomock.Controller) service.AuthService

	mc := gomock.NewController(t)
	defer mc.Finish()

	type in struct {
		ctx context.Context
		dto auth.CheckDTO
	}

	type want struct {
		resp auth.CheckResp
		err  error
	}

	var (
		ctx = context.TODO()

		servErr = errors.New("service error")

		id    = int64(1)
		token = gofakeit.Animal()

		dto = auth.CheckDTO{
			ID:    id,
			Token: token,
		}

		resp = auth.CheckResp{
			ID:       id,
			Username: gofakeit.Username(),
			Token:    token,
			ExpAt:    time.Now(),
		}
	)

	tests := []struct {
		name                string
		in                  in
		want                want
		authServiceMockFunc authServiceMockFunc
	}{
		{
			name: "ok",
			in: in{
				ctx: ctx,
				dto: dto,
			},
			want: want{
				resp: resp,
				err:  nil,
			},
			authServiceMockFunc: func(mc *gomock.Controller) service.AuthService {
				mock := servMocks.NewMockAuthService(mc)
				mock.EXPECT().Check(ctx, dto).Return(resp, nil)
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
				resp: resp,
				err:  servErr,
			},
			authServiceMockFunc: func(mc *gomock.Controller) service.AuthService {
				mock := servMocks.NewMockAuthService(mc)
				mock.EXPECT().Check(ctx, dto).Return(resp, servErr)
				return mock
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			asm := test.authServiceMockFunc(mc)

			r, err := asm.Check(test.in.ctx, test.in.dto)

			if !errors.Is(err, test.want.err) {
				t.Fatalf("got %v; want %v", err, test.want.err)
			}
			if !reflect.DeepEqual(r, test.want.resp) {
				t.Fatalf("got user %v; want %v", r, test.want.resp)
			}
		})
	}
}
