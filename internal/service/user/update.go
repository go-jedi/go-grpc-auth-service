package user

import (
	"context"

	"github.com/go-jedi/auth/internal/domain/user"
)

func (s *serv) Update(ctx context.Context, dto user.UpdateDTO) (user.User, error) {
	return s.userRepository.Update(ctx, dto)
}
