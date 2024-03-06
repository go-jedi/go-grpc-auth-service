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

func TestLogin(t *testing.T) {
	t.Parallel()
	// Arrange
	type authServiceMockFunc func(mc *gomock.Controller) service.AuthService

	mc := gomock.NewController(t)
	defer mc.Finish()

	type input struct {
		ctx          context.Context
		loginRequest *model.LoginRequest
	}

	var (
		ctx = context.Background()

		username = gofakeit.Animal()
		password = gofakeit.Password(true, true, true, true, false, 32)

		accessToken  = gofakeit.Password(true, true, true, true, false, 50)
		refreshToken = gofakeit.Password(true, true, true, true, false, 50)

		serviceErr = fmt.Errorf("service error")

		loginRequest = &model.LoginRequest{
			Username: username,
			Password: password,
		}

		expected = &model.LoginResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}
	)

	tests := []struct {
		name                string
		input               input
		expected            *model.LoginResponse
		err                 error
		authServiceMockFunc authServiceMockFunc
	}{
		{
			name: "OK (Login)",
			input: input{
				ctx:          ctx,
				loginRequest: loginRequest,
			},
			expected: expected,
			err:      nil,
			authServiceMockFunc: func(mc *gomock.Controller) service.AuthService {
				mock := serviceMocks.NewMockAuthService(mc)
				mock.EXPECT().Login(ctx, loginRequest).Return(expected, nil)
				return mock
			},
		},
		{
			name: "Service error case",
			input: input{
				ctx:          ctx,
				loginRequest: loginRequest,
			},
			expected: nil,
			err:      serviceErr,
			authServiceMockFunc: func(mc *gomock.Controller) service.AuthService {
				mock := serviceMocks.NewMockAuthService(mc)
				mock.EXPECT().Login(ctx, loginRequest).Return(nil, serviceErr)
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
			result, err := authServiceMock.Login(test.input.ctx, test.input.loginRequest)

			//	Assert
			require.Equal(t, test.err, err)
			require.Equal(t, test.expected, result)
		})
	}
}
