package auth

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-jedi/platform_common/pkg/db"
	"go.uber.org/zap"

	"github.com/go-jedi/auth-service/internal/logger"
	"github.com/go-jedi/auth-service/internal/model"
	"github.com/go-jedi/auth-service/internal/repository"
	"github.com/go-jedi/auth-service/internal/repository/auth/converter"

	modelRepo "github.com/go-jedi/auth-service/internal/repository/auth/model"
)

const (
	tableName = "users"

	idColumn             = "id"
	usernameColumn       = "username"
	passwordColumn       = "password"
	createdAtColumn      = "created_at"
	updatedAtColumn      = "updated_at"
	passwordLastChangeAt = "password_last_change_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.AuthRepository {
	return &repo{db: db}
}

func (r *repo) Register(ctx context.Context, registerRequest *model.RegisterRequest) error {
	logger.Info(
		"(REPOSITORY) Register auth...",
		zap.String("username", registerRequest.Username),
		zap.String("password", registerRequest.Password),
	)

	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(usernameColumn, passwordColumn).
		Values(registerRequest.Username, registerRequest.Password)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "auth_repository.Register",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	logger.Info(
		"(REPOSITORY) GetUserByUsername auth...",
		zap.String("username", username),
	)

	builder := sq.Select(idColumn, usernameColumn, passwordColumn, createdAtColumn, updatedAtColumn, passwordLastChangeAt).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{usernameColumn: username}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "auth_repository.Login",
		QueryRaw: query,
	}

	var user modelRepo.User
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.PasswordLastChangeAt,
	)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}
