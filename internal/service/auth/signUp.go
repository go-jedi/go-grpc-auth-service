package auth

import (
	"context"
	"log"
	"strconv"

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

	u, err := s.userRepository.Create(ctx, dto)
	if err != nil {
		return user.User{}, err
	}

	// add new user in cache
	if err := s.signUpAddToCache(ctx, u); err != nil {
		log.Println(err)
	}

	return u, nil
}

// signUpAddToCache add cache val for method SignUp.
func (s *serv) signUpAddToCache(ctx context.Context, u user.User) error {
	key := strconv.FormatInt(u.ID, 10)

	return s.cache.User.Set(ctx, key, u, 0)
}
