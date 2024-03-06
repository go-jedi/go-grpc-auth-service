package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-jedi/platform_common/pkg/db"
	"go.uber.org/zap"

	"github.com/go-jedi/auth-service/internal/logger"
	"github.com/go-jedi/auth-service/internal/model"
	"github.com/go-jedi/auth-service/internal/repository"
	"github.com/go-jedi/auth-service/internal/repository/user/converter"

	modelRepo "github.com/go-jedi/auth-service/internal/repository/user/model"
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

func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	logger.Info(
		"(REPOSITORY) Get user...",
		zap.Int64("id", id),
	)

	builder := sq.Select(idColumn, usernameColumn, passwordColumn, createdAtColumn, updatedAtColumn, passwordLastChangeAt).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.Get",
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

func (r *repo) UpdateName(ctx context.Context, updateNameRequest *model.UpdateNameRequest) error {
	logger.Info(
		"(REPOSITORY) UpdateName user...",
		zap.Int64("id", updateNameRequest.ID),
		zap.String("username", updateNameRequest.Username),
	)

	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(usernameColumn, updateNameRequest.Username).
		Where(sq.Eq{idColumn: updateNameRequest.ID})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.UpdateName",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) UpdatePassword(ctx context.Context, updatePasswordRequest *model.UpdatePasswordRequest) error {
	logger.Info(
		"(REPOSITORY) UpdatePassword user...",
		zap.Int64("id", updatePasswordRequest.ID),
		zap.String("password", updatePasswordRequest.Password),
	)

	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(passwordColumn, updatePasswordRequest.Password).
		Where(sq.Eq{idColumn: updatePasswordRequest.ID})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.UpdatePassword",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
