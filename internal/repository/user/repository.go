package user

import (
	"github.com/go-jedi/auth/internal/repository"
	"github.com/go-jedi/auth/pkg/logger"
	"github.com/go-jedi/auth/pkg/postgres"
)

type repo struct {
	logger *logger.Logger
	db     *postgres.Postgres
}

func NewRepository(
	l *logger.Logger,
	p *postgres.Postgres,
) repository.UserRepository {
	return &repo{
		logger: l,
		db:     p,
	}
}
