package user

import (
	"context"
	"log"
	"strconv"

	"github.com/go-jedi/auth/internal/domain/user"
	"github.com/go-jedi/auth/pkg/apperrors"
	"github.com/go-jedi/auth/pkg/bcrypt"
)

func (s *serv) Create(ctx context.Context, dto user.CreateDTO) (user.User, error) {
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
	if err := s.createAddToCache(ctx, u); err != nil {
		log.Println(err)
	}

	return u, nil
}

// createAddToCache add cache val for method Create.
func (s *serv) createAddToCache(ctx context.Context, u user.User) error {
	key := strconv.FormatInt(u.ID, 10)

	return s.cache.User.Set(ctx, key, u, 0)
}
