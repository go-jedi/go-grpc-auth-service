package jwt

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/go-jedi/platform_common/pkg/sys"
	"github.com/go-jedi/platform_common/pkg/sys/codes"
	"google.golang.org/grpc/metadata"

	"github.com/pkg/errors"

	"github.com/go-jedi/auth-service/internal/model"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(id int64, secretKey []byte, duration time.Duration) (string, error) {
	claims := model.UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
		ID: id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func VerifyToken(tokenStr string, secretKey []byte) (*model.UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&model.UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.Errorf("unexpected token signing method")
			}

			return secretKey, nil
		},
	)
	if err != nil {
		return nil, errors.Errorf("invalid token: %s", err.Error())
	}

	claims, ok := token.Claims.(*model.UserClaims)
	if !ok {
		return nil, errors.Errorf("invalid token claims")
	}

	return claims, nil
}

func Check(ctx context.Context) (bool, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false, sys.NewCommonError("metadata is not provided", codes.Unauthenticated)
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return false, sys.NewCommonError("authorization header is not provided", codes.Unauthenticated)
	}

	if !strings.HasPrefix(authHeader[0], "Bearer ") {
		return false, sys.NewCommonError("invalid authorization header format", codes.Unauthenticated)
	}

	accessToken := strings.TrimPrefix(authHeader[0], "Bearer ")

	_, err := VerifyToken(accessToken, []byte(os.Getenv("ACCESS_TOKEN_SECRET_KEY")))
	if err != nil {
		return false, sys.NewCommonError("invalid access token", codes.Unauthenticated)
	}

	return true, nil
}
