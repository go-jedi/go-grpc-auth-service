package auth

import (
	"context"

	"github.com/go-jedi/auth/internal/domain/auth"
	"github.com/go-jedi/auth/pkg/apperrors"
	"github.com/go-jedi/auth/pkg/bcrypt"
)

func (s *serv) SignIn(ctx context.Context, dto auth.SignInDTO) (auth.SignInResp, error) {
	// get user by username
	u, err := s.userRepository.GetByUsername(ctx, dto.Username)
	if err != nil {
		return auth.SignInResp{}, err
	}

	// compare password hash and password
	if err := bcrypt.CompareHashAndPassword(u.PasswordHash, dto.Password); err != nil {
		return auth.SignInResp{}, apperrors.ErrUserIncorrectPassword
	}

	// generate access, refresh tokens
	t, err := s.jwt.Generate(u.ID, u.Username)
	if err != nil {
		return auth.SignInResp{}, err
	}

	return t, nil
}
