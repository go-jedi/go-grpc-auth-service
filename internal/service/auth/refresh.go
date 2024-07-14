package auth

import (
	"context"

	"github.com/go-jedi/auth/internal/domain/auth"
)

func (s *serv) Refresh(_ context.Context, _ auth.RefreshDTO) (auth.RefreshResp, error) {
	return auth.RefreshResp{}, nil
}
