package auth

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/go-jedi/platform_common/pkg/sys"
	"github.com/go-jedi/platform_common/pkg/sys/codes"

	"github.com/go-jedi/auth-service/internal/utils/jwt"

	"github.com/go-jedi/auth-service/internal/logger"
	"go.uber.org/zap"
)

func (s *serv) GetRefreshToken(_ context.Context, refreshToken string) (string, error) {
	logger.Info(
		"(SERVICE) GetRefreshToken auth...",
		zap.String("refreshToken", refreshToken),
	)

	claims, err := jwt.VerifyToken(refreshToken, []byte(os.Getenv("REFRESH_TOKEN_SECRET_KEY")))
	if err != nil {
		return "", sys.NewCommonError("invalid refresh token", codes.Unauthenticated)
	}

	refreshTokenDuration, err := strconv.ParseInt(os.Getenv("REFRESH_TOKEN_EXPIRATION"), 10, 64)
	if err != nil {
		return "", err
	}

	refreshToken, err = jwt.GenerateToken(
		claims.ID,
		[]byte(os.Getenv("REFRESH_TOKEN_SECRET_KEY")),
		time.Duration(refreshTokenDuration),
	)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
