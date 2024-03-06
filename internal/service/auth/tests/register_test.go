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

func TestRegister(t *testing.T) {
	t.Parallel()
	// Arrange
	type authServiceMockFunc func(mc *gomock.Controller) service.AuthService

	mc := gomock.NewController(t)
	defer mc.Finish()

	type input struct {
		ctx             context.Context
		registerRequest *model.RegisterRequest
	}

	var (
		ctx = context.Background()

		username = gofakeit.Animal()
		password = gofakeit.Password(true, true, true, true, false, 32)

		serviceErr = fmt.Errorf("service error")

		registerRequest = &model.RegisterRequest{
			Username: username,
			Password: password,
		}
	)

	tests := []struct {
		name                string
		input               input
		expected            error
		authServiceMockFunc authServiceMockFunc
	}{
		{
			name: "OK (Register)",
			input: input{
				ctx:             ctx,
				registerRequest: registerRequest,
			},
			expected: nil,
			authServiceMockFunc: func(mc *gomock.Controller) service.AuthService {
				mock := serviceMocks.NewMockAuthService(mc)
				mock.EXPECT().Register(ctx, registerRequest).Return(nil)
				return mock
			},
		},
		{
			name: "Service error case",
			input: input{
				ctx:             ctx,
				registerRequest: registerRequest,
			},
			expected: serviceErr,
			authServiceMockFunc: func(mc *gomock.Controller) service.AuthService {
				mock := serviceMocks.NewMockAuthService(mc)
				mock.EXPECT().Register(ctx, registerRequest).Return(serviceErr)
				return mock
			},
		},
	}
	//	Act
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			authServiceMock := test.authServiceMockFunc(mc)
			err := authServiceMock.Register(test.input.ctx, test.input.registerRequest)

			//	Assert
			require.Equal(t, test.expected, err)
		})
	}
}
