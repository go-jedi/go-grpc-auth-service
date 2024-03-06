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

func TestUpdateName(t *testing.T) {
	t.Parallel()
	// Arrange
	type userServiceMockFunc func(mc *gomock.Controller) service.UserService

	mc := gomock.NewController(t)
	defer mc.Finish()

	type input struct {
		ctx               context.Context
		updateNameRequest *model.UpdateNameRequest
	}

	var (
		ctx = context.Background()

		id       = gofakeit.Int64()
		username = gofakeit.Animal()

		serviceErr = fmt.Errorf("service error")

		updateNameRequest = &model.UpdateNameRequest{
			ID:       id,
			Username: username,
		}
	)

	tests := []struct {
		name                string
		input               input
		expected            error
		userServiceMockFunc userServiceMockFunc
	}{
		{
			name: "OK (UpdateName)",
			input: input{
				ctx:               ctx,
				updateNameRequest: updateNameRequest,
			},
			expected: nil,
			userServiceMockFunc: func(mc *gomock.Controller) service.UserService {
				mock := serviceMocks.NewMockUserService(mc)
				mock.EXPECT().UpdateName(ctx, updateNameRequest).Return(nil)
				return mock
			},
		},
		{
			name: "Service error case",
			input: input{
				ctx:               ctx,
				updateNameRequest: updateNameRequest,
			},
			expected: serviceErr,
			userServiceMockFunc: func(mc *gomock.Controller) service.UserService {
				mock := serviceMocks.NewMockUserService(mc)
				mock.EXPECT().UpdateName(ctx, updateNameRequest).Return(serviceErr)
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
			err := userServiceMock.UpdateName(test.input.ctx, test.input.updateNameRequest)

			//	Assert
			require.Equal(t, test.expected, err)
		})
	}
}
