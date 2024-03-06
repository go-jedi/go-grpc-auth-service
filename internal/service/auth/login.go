package auth

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/go-jedi/auth-service/internal/utils/jwt"

	"github.com/go-jedi/platform_common/pkg/sys"
	"github.com/go-jedi/platform_common/pkg/sys/codes"

	"github.com/go-jedi/auth-service/internal/utils/bcrypt"

	"github.com/go-jedi/auth-service/internal/logger"
	"github.com/go-jedi/auth-service/internal/model"
	"go.uber.org/zap"
)

func (s *serv) Login(ctx context.Context, loginRequest *model.LoginRequest) (*model.LoginResponse, error) {
	logger.Info(
		"(SERVICE) Login auth...",
		zap.String("username", loginRequest.Username),
		zap.String("password", loginRequest.Password),
	)

	user, err := s.authRepository.GetUserByUsername(ctx, loginRequest.Username)
	if err != nil {
		return nil, err
	}

	result := bcrypt.VerifyPassword(user.Password, loginRequest.Password)
	if !result {
		return nil, sys.NewCommonError("Неправильный логин или пароль", codes.InvalidArgument)
	}

	accessTokenDuration, err := strconv.ParseInt(os.Getenv("ACCESS_TOKEN_EXPIRATION"), 10, 64)
	if err != nil {
		return nil, err
	}

	refreshTokenDuration, err := strconv.ParseInt(os.Getenv("REFRESH_TOKEN_EXPIRATION"), 10, 64)
	if err != nil {
		return nil, err
	}

	accessToken, err := jwt.GenerateToken(
		user.ID,
		[]byte(os.Getenv("ACCESS_TOKEN_SECRET_KEY")),
		time.Duration(accessTokenDuration),
	)
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.GenerateToken(
		user.ID,
		[]byte(os.Getenv("REFRESH_TOKEN_SECRET_KEY")),
		time.Duration(refreshTokenDuration),
	)
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
