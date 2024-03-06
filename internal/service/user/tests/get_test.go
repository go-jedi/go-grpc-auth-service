package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/go-jedi/auth-service/internal/model"
	"github.com/go-jedi/auth-service/internal/service"

	serviceMocks "github.com/go-jedi/auth-service/internal/service/mocks"
)

func TestGet(t *testing.T) {
	t.Parallel()
	// Arrange
	type userServiceMockFunc func(mc *gomock.Controller) service.UserService

	mc := gomock.NewController(t)
	defer mc.Finish()

	type input struct {
		ctx context.Context
		id  int64
	}

	var (
		ctx = context.Background()

		id       = gofakeit.Int64()
		username = gofakeit.Animal()
		password = gofakeit.Password(true, true, true, true, false, 32)

		serviceErr = fmt.Errorf("service error")

		expected = &model.User{
			ID:                   id,
			Username:             username,
			Password:             password,
			CreatedAt:            time.Now(),
			UpdatedAt:            sql.NullTime{},
			PasswordLastChangeAt: time.Now(),
		}
	)

	tests := []struct {
		name                string
		input               input
		expected            *model.User
		err                 error
		userServiceMockFunc userServiceMockFunc
	}{
		{
			name: "OK (Get)",
			input: input{
				ctx: ctx,
				id:  id,
			},
			expected: expected,
			err:      nil,
			userServiceMockFunc: func(mc *gomock.Controller) service.UserService {
				mock := serviceMocks.NewMockUserService(mc)
				mock.EXPECT().Get(ctx, id).Return(expected, nil)
				return mock
			},
		},
		{
			name: "Service error case",
			input: input{
				ctx: ctx,
				id:  id,
			},
			expected: nil,
			err:      serviceErr,
			userServiceMockFunc: func(mc *gomock.Controller) service.UserService {
				mock := serviceMocks.NewMockUserService(mc)
				mock.EXPECT().Get(ctx, id).Return(nil, serviceErr)
				return mock
			},
		},
	}
	//	Act
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			userServiceMock := test.userServiceMockFunc(mc)
			result, err := userServiceMock.Get(test.input.ctx, test.input.id)

			//	Assert
			require.Equal(t, test.err, err)
			require.Equal(t, test.expected, result)
		})
	}
}
