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

func TestUpdateName(t *testing.T) {
	t.Parallel()
	//	 Arrange
	type userRepositoryMockFunc func(mc *gomock.Controller) repository.UserRepository

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

		repoErr = fmt.Errorf("repository error")

		updateNameRequest = &model.UpdateNameRequest{
			ID:       id,
			Username: username,
		}
	)

	tests := []struct {
		name                   string
		input                  input
		expected               error
		userRepositoryMockFunc userRepositoryMockFunc
	}{
		{
			name: "OK (UpdateName)",
			input: input{
				ctx:               ctx,
				updateNameRequest: updateNameRequest,
			},
			expected: nil,
			userRepositoryMockFunc: func(mc *gomock.Controller) repository.UserRepository {
				mock := repoMocks.NewMockUserRepository(mc)
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
			expected: repoErr,
			userRepositoryMockFunc: func(*gomock.Controller) repository.UserRepository {
				mock := repoMocks.NewMockUserRepository(mc)
				mock.EXPECT().UpdateName(ctx, updateNameRequest).Return(repoErr)
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
			err := userRepositoryMock.UpdateName(test.input.ctx, test.input.updateNameRequest)

			//	Assert
			require.Equal(t, test.expected, err)
		})
	}
}
