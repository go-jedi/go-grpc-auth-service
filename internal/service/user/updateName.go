package user

import (
	"context"

	"github.com/go-jedi/auth-service/internal/logger"
	"go.uber.org/zap"

	"github.com/go-jedi/auth-service/internal/model"
)

func (s *serv) UpdateName(ctx context.Context, updateNameRequest *model.UpdateNameRequest) error {
	logger.Info(
		"(SERVICE) UpdateName user...",
		zap.Int64("id", updateNameRequest.ID),
		zap.String("username", updateNameRequest.Username),
	)

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		err := s.userRepository.UpdateName(ctx, updateNameRequest)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
