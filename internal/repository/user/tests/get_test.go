package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-jedi/auth-service/internal/model"

	"github.com/go-jedi/auth-service/internal/repository"
	"github.com/golang/mock/gomock"

	repoMocks "github.com/go-jedi/auth-service/internal/repository/mocks"
)

func TestGet(t *testing.T) {
	t.Parallel()
	//	 Arrange
	type userRepositoryMockFunc func(mc *gomock.Controller) repository.UserRepository

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

		repoErr = fmt.Errorf("repository error")

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
		name                   string
		input                  input
		expected               *model.User
		err                    error
		userRepositoryMockFunc userRepositoryMockFunc
	}{
		{
			name: "OK (Get)",
			input: input{
				ctx: ctx,
				id:  id,
			},
			expected: expected,
			err:      nil,
			userRepositoryMockFunc: func(mc *gomock.Controller) repository.UserRepository {
				mock := repoMocks.NewMockUserRepository(mc)
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
			err:      repoErr,
			userRepositoryMockFunc: func(mc *gomock.Controller) repository.UserRepository {
				mock := repoMocks.NewMockUserRepository(mc)
				mock.EXPECT().Get(ctx, id).Return(nil, repoErr)
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
			result, err := userRepositoryMock.Get(test.input.ctx, test.input.id)

			//	Assert
			require.Equal(t, test.err, err)
			require.Equal(t, test.expected, result)
		})
	}
}
