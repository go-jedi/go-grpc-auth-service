package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/go-jedi/auth-service/internal/model"
	"github.com/go-jedi/auth-service/internal/service"

	serviceMocks "github.com/go-jedi/auth-service/internal/service/mocks"
)

func TestUpdatePassword(t *testing.T) {
	t.Parallel()
	// Arrange
	type userServiceMockFunc func(mc *gomock.Controller) service.UserService

	mc := gomock.NewController(t)
	defer mc.Finish()

	type input struct {
		ctx                   context.Context
		updatePasswordRequest *model.UpdatePasswordRequest
	}

	var (
		ctx = context.Background()

		id       = gofakeit.Int64()
		password = gofakeit.Password(true, true, true, true, false, 32)

		serviceErr = fmt.Errorf("service error")

		updatePasswordRequest = &model.UpdatePasswordRequest{
			ID:       id,
			Password: password,
		}
	)

	tests := []struct {
		name                string
		input               input
		expected            error
		userServiceMockFunc userServiceMockFunc
	}{
		{
			name: "OK (UpdatePassword)",
			input: input{
				ctx:                   ctx,
				updatePasswordRequest: updatePasswordRequest,
			},
			expected: nil,
			userServiceMockFunc: func(mc *gomock.Controller) service.UserService {
				mock := serviceMocks.NewMockUserService(mc)
				mock.EXPECT().UpdatePassword(ctx, updatePasswordRequest).Return(nil)
				return mock
			},
		},
		{
			name: "Service error case",
			input: input{
				ctx:                   ctx,
				updatePasswordRequest: updatePasswordRequest,
			},
			expected: serviceErr,
			userServiceMockFunc: func(mc *gomock.Controller) service.UserService {
				mock := serviceMocks.NewMockUserService(mc)
				mock.EXPECT().UpdatePassword(ctx, updatePasswordRequest).Return(serviceErr)
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
			err := userServiceMock.UpdatePassword(test.input.ctx, test.input.updatePasswordRequest)

			//	Assert
			require.Equal(t, test.expected, err)
		})
	}
}
