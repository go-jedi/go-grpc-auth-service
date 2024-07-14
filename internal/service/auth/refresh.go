package auth

import (
	"context"

	"github.com/go-jedi/auth/internal/domain/auth"
)

func (s *serv) Refresh(ctx context.Context, dto auth.RefreshDTO) (auth.RefreshResp, error) {
	u, err := s.userRepository.GetByID(ctx, dto.ID)
	if err != nil {
		return auth.RefreshResp{}, err
	}

	// check verify token
	vr, err := s.jwt.Verify(u.ID, u.Username, dto.RefreshToken)
	if err != nil {
		return auth.RefreshResp{}, err
	}

	// generate access, refresh tokens
	gr, err := s.jwt.Generate(vr.ID, vr.Username)
	if err != nil {
		return auth.RefreshResp{}, err
	}

	return auth.RefreshResp{
		AccessToken:  gr.AccessToken,
		RefreshToken: gr.RefreshToken,
		AccessExpAt:  gr.AccessExpAt,
		RefreshExpAt: gr.RefreshExpAt,
	}, nil
}
