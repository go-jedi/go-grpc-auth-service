package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/go-jedi/auth-service/internal/model"
	"github.com/go-jedi/auth-service/internal/repository"
	"github.com/golang/mock/gomock"

	repoMocks "github.com/go-jedi/auth-service/internal/repository/mocks"
)

func TestRegister(t *testing.T) {
	t.Parallel()
	// Arrange
	type authRepositoryMockFunc func(mc *gomock.Controller) repository.AuthRepository

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

		repoErr = fmt.Errorf("repository error")

		registerRequest = &model.RegisterRequest{
			Username: username,
			Password: password,
		}
	)

	tests := []struct {
		name                   string
		input                  input
		expected               error
		authRepositoryMockFunc authRepositoryMockFunc
	}{
		{
			name: "OK (Register)",
			input: input{
				ctx:             ctx,
				registerRequest: registerRequest,
			},
			expected: nil,
			authRepositoryMockFunc: func(mc *gomock.Controller) repository.AuthRepository {
				mock := repoMocks.NewMockAuthRepository(mc)
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
			expected: repoErr,
			authRepositoryMockFunc: func(mc *gomock.Controller) repository.AuthRepository {
				mock := repoMocks.NewMockAuthRepository(mc)
				mock.EXPECT().Register(ctx, registerRequest).Return(repoErr)
				return mock
			},
		},
	}
	//	Act
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			authRepositoryMock := test.authRepositoryMockFunc(mc)
			err := authRepositoryMock.Register(test.input.ctx, test.input.registerRequest)

			//	Assert
			require.Equal(t, test.expected, err)
		})
	}
}
