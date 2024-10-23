package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPwd return the bcrypt hash of the password
func HashPwd(pwd string) (string, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPwd), nil
}

// CheckPwd checks if the provided password is correct or not
func CheckPwd(pwd string, hashedPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
}
