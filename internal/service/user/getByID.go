package user

import (
	"context"

	"github.com/go-jedi/auth/internal/domain/user"
)

func (s *serv) GetByID(ctx context.Context, id int64) (user.User, error) {
	s.logger.Info("SERVICE: GetByID")

	return s.userRepository.GetByID(ctx, id)
}
