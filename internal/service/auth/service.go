package auth

import (
	"github.com/go-jedi/auth-service/internal/repository"
	"github.com/go-jedi/auth-service/internal/service"
	"github.com/go-jedi/platform_common/pkg/db"
)

type serv struct {
	authRepository repository.AuthRepository
	txManager      db.TxManager
}

func NewService(
	authRepository repository.AuthRepository,
	txManager db.TxManager,
) service.AuthService {
	return &serv{
		authRepository: authRepository,
		txManager:      txManager,
	}
}
