package user

import (
	"context"

	"github.com/go-jedi/auth/internal/domain/user"
	"github.com/go-jedi/auth/pkg/apperrors"
	"github.com/go-jedi/auth/pkg/bcrypt"
)

func (s *serv) Update(ctx context.Context, dto user.UpdateDTO) (user.User, error) {
	// check user exists
	if _, err := s.userRepository.GetByID(ctx, dto.ID); err != nil {
		return user.User{}, err
	}

	// TODO: check exist user new data here...

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
