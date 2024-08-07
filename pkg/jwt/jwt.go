package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/go-jedi/auth/config"
	"github.com/go-jedi/auth/pkg/uid"
	"github.com/golang-jwt/jwt/v5"
)

const (
	defaultSecretHashLen = 20
	defaultAccessExpAt   = 5
	defaultRefreshExpAt  = 30
)

var (
	ErrTokenSigningMethod = errors.New("unexpected token signing method")
	ErrTokenInvalid       = errors.New("invalid token")
	ErrTokenClaims        = errors.New("unexpected token claims")
	ErrTokenExpired       = errors.New("token has expired")
	ErrTokenID            = errors.New("unexpected token id")
	ErrTokenUsername      = errors.New("unexpected token username")
)

type tokenClaims struct {
	ID       int64
	Username string
	jwt.RegisteredClaims
}

type JWT struct {
	// secret key need for token signing
	secret []byte
	// secretHashLen need to generate hash
	secretHashLen int
	// accessExpAt expiration time in minutes
	accessExpAt int
	// refreshExpAt expiration time in days
	refreshExpAt int
}

func (j *JWT) init() error {
	if j.secretHashLen == 0 {
		j.secretHashLen = defaultSecretHashLen
	}

	if j.accessExpAt == 0 {
		j.accessExpAt = defaultAccessExpAt
	}

	if j.refreshExpAt == 0 {
		j.refreshExpAt = defaultRefreshExpAt
	}

	return nil
}

func NewJWT(cfg config.JWTConfig) (*JWT, error) {
	j := &JWT{
		secretHashLen: cfg.SecretHashLen,
		accessExpAt:   cfg.AccessExpAt,
		refreshExpAt:  cfg.RefreshExpAt,
	}

	if err := j.generateSecretKey(cfg.SecretPath); err != nil {
		return nil, err
	}

	return j, nil
}

type GenerateResp struct {
	AccessToken  string
	RefreshToken string
	AccessExpAt  time.Time
	RefreshExpAt time.Time
}

// Generate token.
func (j *JWT) Generate(id int64, username string) (GenerateResp, error) {
	aExpAt := j.getAccessExpAt()
	rExpAt := j.getRefreshExpAt()

	aToken, err := j.createToken(id, username, aExpAt)
	if err != nil {
		return GenerateResp{}, err
	}

	rToken, err := j.createToken(id, username, rExpAt)
	if err != nil {
		return GenerateResp{}, err
	}

	return GenerateResp{
		AccessToken:  aToken,
		RefreshToken: rToken,
		AccessExpAt:  aExpAt,
		RefreshExpAt: rExpAt,
	}, nil
}

type VerifyResp struct {
	ID       int64
	Username string
	ExpAt    time.Time
}

// Verify token.
func (j *JWT) Verify(id int64, username string, token string) (VerifyResp, error) {
	// parse the token
	t, err := jwt.ParseWithClaims(
		token,
		&tokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, ErrTokenSigningMethod
			}
			return j.secret, nil
		})
	if err != nil {
		return VerifyResp{}, err
	}

	// check token valid
	if !t.Valid {
		return VerifyResp{}, ErrTokenInvalid
	}

	// extract the claims
	c, ok := t.Claims.(*tokenClaims)
	if !ok {
		return VerifyResp{}, ErrTokenClaims
	}

	// check expired token
	if c.ExpiresAt != nil && time.Now().After(c.ExpiresAt.Time) {
		return VerifyResp{}, ErrTokenExpired
	}

	// compare id with id in token
	if id != c.ID {
		return VerifyResp{}, ErrTokenID
	}

	// compare username with username in token
	if username != c.Username {
		return VerifyResp{}, ErrTokenUsername
	}

	return VerifyResp{
		ID:       c.ID,
		Username: c.Username,
		ExpAt:    c.ExpiresAt.Time,
	}, nil
}

// createToken create token.
func (j *JWT) createToken(id int64, username string, expAt time.Time) (token string, err error) {
	// create the claims
	c := tokenClaims{
		ID:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// create token
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(j.secret)
	if err != nil {
		return
	}

	return
}

// generateSecretKey generate secret key.
func (j *JWT) generateSecretKey(sp string) error {
	ie, err := j.fileExists(sp)
	if err != nil {
		return err
	}

	if ie {
		fb, err := os.ReadFile(sp)
		if err != nil {
			return err
		}
		j.secret = fb
		return nil
	}

	u, err := uid.GenSpecialUID(j.secretHashLen)
	if err != nil {
		return err
	}

	j.secret = []byte(u)

	if err := os.WriteFile(sp, j.secret, 0666); err != nil {
		return err
	}

	return nil
}

// getAccessExpAt get access expires at token time.
func (j *JWT) getAccessExpAt() time.Time {
	return time.Now().Add(time.Duration(j.accessExpAt) * time.Minute)
}

// getRefreshExpAt get refresh expires at token time.
func (j *JWT) getRefreshExpAt() time.Time {
	return time.Now().Add(time.Duration(j.refreshExpAt) * 24 * time.Hour)
}

//
// UTILS
//

// fileExists check file exists.
func (j *JWT) fileExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return !fi.IsDir(), nil
}
