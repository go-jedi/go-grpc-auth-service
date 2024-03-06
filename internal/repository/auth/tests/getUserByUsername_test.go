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

func TestGetUserByUsername(t *testing.T) {
	t.Parallel()

	// Arrange
	type authRepositoryMockFunc func(mc *gomock.Controller) repository.AuthRepository

	mc := gomock.NewController(t)
	defer mc.Finish()

	type input struct {
		ctx      context.Context
		username string
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
		authRepositoryMockFunc authRepositoryMockFunc
	}{
		{
			name: "OK (GetUserByUsername)",
			input: input{
				ctx:      ctx,
				username: username,
			},
			expected: expected,
			err:      nil,
			authRepositoryMockFunc: func(mc *gomock.Controller) repository.AuthRepository {
				mock := repoMocks.NewMockAuthRepository(mc)
				mock.EXPECT().GetUserByUsername(ctx, username).Return(expected, nil)
				return mock
			},
		},
		{
			name: "Service error case",
			input: input{
				ctx:      ctx,
				username: username,
			},
			expected: nil,
			err:      repoErr,
			authRepositoryMockFunc: func(mc *gomock.Controller) repository.AuthRepository {
				mock := repoMocks.NewMockAuthRepository(mc)
				mock.EXPECT().GetUserByUsername(ctx, username).Return(nil, repoErr)
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
			result, err := authRepositoryMock.GetUserByUsername(test.input.ctx, test.input.username)

			//	Assert
			require.Equal(t, test.err, err)
			require.Equal(t, test.expected, result)
		})
	}
}
