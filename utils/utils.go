package utils

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateUuid() string {
	uuid := uuid.New()
	return uuid.String()
}

func GeneratePassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
