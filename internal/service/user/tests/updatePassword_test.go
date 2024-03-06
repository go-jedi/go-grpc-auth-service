package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	repoMocks "github.com/go-jedi/auth-service/internal/repository/mocks"
	"github.com/stretchr/testify/require"

	"github.com/go-jedi/auth-service/internal/model"
	"github.com/go-jedi/auth-service/internal/repository"
	"github.com/golang/mock/gomock"
)

func TestUpdatePassword(t *testing.T) {
	t.Parallel()
	//	 Arrange
	type userRepositoryMockFunc func(mc *gomock.Controller) repository.UserRepository

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

		repoErr = fmt.Errorf("repository error")

		updatePasswordRequest = &model.UpdatePasswordRequest{
			ID:       id,
			Password: password,
		}
	)

	tests := []struct {
		name                   string
		input                  input
		expected               error
		userRepositoryMockFunc userRepositoryMockFunc
	}{
		{
			name: "OK (UpdatePassword)",
			input: input{
				ctx:                   ctx,
				updatePasswordRequest: updatePasswordRequest,
			},
			expected: nil,
			userRepositoryMockFunc: func(mc *gomock.Controller) repository.UserRepository {
				mock := repoMocks.NewMockUserRepository(mc)
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
			expected: repoErr,
			userRepositoryMockFunc: func(mc *gomock.Controller) repository.UserRepository {
				mock := repoMocks.NewMockUserRepository(mc)
				mock.EXPECT().UpdatePassword(ctx, updatePasswordRequest).Return(repoErr)
				return mock
			},
		},
	}
	//	Act
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			userRepositoryMock := test.userRepositoryMockFunc(mc)
			err := userRepositoryMock.UpdatePassword(test.input.ctx, test.input.updatePasswordRequest)

			//	Assert
			require.Equal(t, test.expected, err)
		})
	}
}
