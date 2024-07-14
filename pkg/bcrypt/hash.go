package bcrypt

import "golang.org/x/crypto/bcrypt"

// GenerateHash generate hash.
func GenerateHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
