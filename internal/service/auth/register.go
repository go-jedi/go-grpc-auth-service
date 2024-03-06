package auth

import (
	"context"

	"github.com/go-jedi/auth-service/internal/logger"
	"github.com/go-jedi/auth-service/internal/model"
	"github.com/go-jedi/auth-service/internal/utils/bcrypt"
	"go.uber.org/zap"
)

func (s *serv) Register(ctx context.Context, registerRequest *model.RegisterRequest) error {
	logger.Info(
		"(SERVICE) Register auth...",
		zap.String("username", registerRequest.Username),
		zap.String("password", registerRequest.Password),
	)

	result, err := bcrypt.HashPassword(registerRequest.Password)
	if err != nil {
		return err
	}

	registerRequest.Password = result

	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		err = s.authRepository.Register(ctx, registerRequest)
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
