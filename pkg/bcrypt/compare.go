package bcrypt

import "golang.org/x/crypto/bcrypt"

// CompareHashAndPassword compare hash password and password.
func CompareHashAndPassword(hashPassword string, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password)); err != nil {
		return err
	}
	return nil
}
