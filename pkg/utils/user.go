package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

func ComparePasswords(hashedPassword string, password []byte) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), password)
}
