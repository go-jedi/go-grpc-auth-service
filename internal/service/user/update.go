package user

import (
	"context"
	"log"
	"strconv"

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

	u, err = s.userRepository.Update(ctx, dto)
	if err != nil {
		return user.User{}, err
	}

	// update user in cache
	if err := s.updateCache(ctx, u); err != nil {
		log.Println(err)
	}

	return u, err
}

// updateCache update value in cache.
func (s *serv) updateCache(ctx context.Context, u user.User) error {
	key := strconv.FormatInt(u.ID, 10)

	return s.cache.User.Set(ctx, key, u, 0)
}
