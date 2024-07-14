package auth

import (
	"context"

	"github.com/go-jedi/auth/internal/domain/auth"
)

func (s *serv) Check(ctx context.Context, dto auth.CheckDTO) (auth.CheckResp, error) {
	u, err := s.userRepository.GetByID(ctx, dto.ID)
	if err != nil {
		return auth.CheckResp{}, err
	}

	// check verify token
	vr, err := s.jwt.Verify(u.ID, u.Username, dto.Token)
	if err != nil {
		return auth.CheckResp{}, err
	}

	return auth.CheckResp{
		ID:       vr.ID,
		Username: vr.Username,
		Token:    dto.Token,
		ExpAt:    vr.ExpAt,
	}, nil
}
