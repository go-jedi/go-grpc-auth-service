package user

import (
	"context"

	"github.com/go-jedi/auth/internal/domain/user"
	"github.com/go-jedi/auth/pkg/apperrors"
	"github.com/go-jedi/auth/pkg/bcrypt"
)

func (s *serv) Update(ctx context.Context, dto user.UpdateDTO) (user.User, error) {
	// check user exists
	u, err := s.userRepository.GetByID(ctx, dto.ID)
	if err != nil {
		return user.User{}, err
	}

	// check exists if new username
	if u.Username != dto.Username {
		ie, err := s.userRepository.ExistsUsername(ctx, dto.Username)
		if err != nil {
			return user.User{}, err
		}
		if ie {
			return user.User{}, apperrors.ErrUsernameAlreadyExists
		}
	}

	// check exists if new email
	if u.Email != dto.Email {
		ie, err := s.userRepository.ExistsEmail(ctx, dto.Email)
		if err != nil {
			return user.User{}, err
		}
		if ie {
			return user.User{}, apperrors.ErrEmailAlreadyExists
		}
	}

	// if password is change to generate hash password
	if !bcrypt.IsBcryptHash(dto.Password) {
		// generate password hash
		hp, err := bcrypt.GenerateHash(dto.Password)
		if err != nil {
			return user.User{}, apperrors.ErrUserPasswordNotGenerated
		}
		dto.Password = hp
	}

	return s.userRepository.Update(ctx, dto)
}
