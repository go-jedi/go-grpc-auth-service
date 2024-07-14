package auth

import (
	"context"

	"github.com/go-jedi/auth/internal/domain/user"
	"github.com/go-jedi/auth/pkg/apperrors"
	"github.com/go-jedi/auth/pkg/bcrypt"
)

func (s *serv) SignUp(ctx context.Context, dto user.CreateDTO) (user.User, error) {
	// check user exists
	ie, err := s.userRepository.Exists(ctx, dto.Username, dto.Email)
	if err != nil {
		return user.User{}, err
	}
	if ie {
		return user.User{}, apperrors.ErrUserAlreadyExists
	}

	// generate password hash
	dto.Password, err = bcrypt.GenerateHash(dto.Password)
	if err != nil {
		return user.User{}, apperrors.ErrUserPasswordNotGenerated
	}

	return s.userRepository.Create(ctx, dto)
}
