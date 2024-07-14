package auth

import (
	"context"

	"github.com/go-jedi/auth/internal/domain/auth"
)

func (s *serv) Check(_ context.Context, _ auth.CheckDTO) (auth.CheckResp, error) {
	return auth.CheckResp{}, nil
}
