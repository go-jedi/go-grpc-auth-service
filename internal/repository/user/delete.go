package user

import (
	"context"

	"github.com/go-jedi/auth/pkg/apperrors"
)

func (r *repo) Delete(ctx context.Context, id int64) error {
	q := `DELETE FROM users WHERE id = $1;`

	ct, err := r.db.Pool.Exec(ctx, q, id)
	if err != nil {
		return err
	}

	if ct.RowsAffected() != 1 {
		return apperrors.ErrNoRowsWereUpdate
	}

	return nil
}
