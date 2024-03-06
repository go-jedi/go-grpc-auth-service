package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/go-jedi/auth-service/internal/service"

	serviceMocks "github.com/go-jedi/auth-service/internal/service/mocks"
)

func TestGetAccessToken(t *testing.T) {
	t.Parallel()
	// Arrange
	type authServiceMockFunc func(mc *gomock.Controller) service.AuthService

	mc := gomock.NewController(t)
	defer mc.Finish()

	type input struct {
		ctx          context.Context
		refreshToken string
	}

	var (
		ctx = context.Background()

		refreshToken = gofakeit.Password(true, true, true, true, false, 50)

		serviceErr = fmt.Errorf("service error")

		expected = gofakeit.Password(true, true, true, true, false, 50)
	)

	tests := []struct {
		name                string
		input               input
		expected            string
		err                 error
		authServiceMockFunc authServiceMockFunc
	}{
		{
			name: "OK (GetAccessToken)",
			input: input{
				ctx:          ctx,
				refreshToken: refreshToken,
			},
			expected: expected,
			err:      nil,
			authServiceMockFunc: func(mc *gomock.Controller) service.AuthService {
				mock := serviceMocks.NewMockAuthService(mc)
				mock.EXPECT().GetAccessToken(ctx, refreshToken).Return(expected, nil)
				return mock
			},
		},
		{
			name: "Service error case",
			input: input{
				ctx:          ctx,
				refreshToken: refreshToken,
			},
			expected: "",
			err:      serviceErr,
			authServiceMockFunc: func(mc *gomock.Controller) service.AuthService {
				mock := serviceMocks.NewMockAuthService(mc)
				mock.EXPECT().GetAccessToken(ctx, refreshToken).Return("", serviceErr)
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
			result, err := authServiceMock.GetAccessToken(test.input.ctx, test.input.refreshToken)

			//	Assert
			require.Equal(t, test.err, err)
			require.Equal(t, test.expected, result)
		})
	}
}
