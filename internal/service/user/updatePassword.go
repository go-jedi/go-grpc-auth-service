package user

import (
	"context"

	"github.com/go-jedi/auth-service/internal/utils/bcrypt"

	"github.com/go-jedi/auth-service/internal/logger"
	"go.uber.org/zap"

	"github.com/go-jedi/auth-service/internal/model"
)

func (s *serv) UpdatePassword(ctx context.Context, updatePasswordRequest *model.UpdatePasswordRequest) error {
	logger.Info(
		"(SERVICE) UpdatePassword user...",
		zap.Int64("id", updatePasswordRequest.ID),
		zap.String("password", updatePasswordRequest.Password),
	)

	result, err := bcrypt.HashPassword(updatePasswordRequest.Password)
	if err != nil {
		return err
	}

	updatePasswordRequest.Password = result

	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		err = s.userRepository.UpdatePassword(ctx, updatePasswordRequest)
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
