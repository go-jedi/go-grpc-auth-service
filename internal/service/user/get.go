package user

import (
	"context"

	"github.com/go-jedi/auth-service/internal/logger"
	"go.uber.org/zap"

	"github.com/go-jedi/auth-service/internal/model"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.User, error) {
	logger.Info(
		"(SERVICE) Get user...",
		zap.Int64("id", id),
	)

	result, err := s.userRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}
