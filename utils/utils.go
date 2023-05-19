package utils

import (
	"github.com/google/uuid"
)

func GenerateUuid() string {
	uuid := uuid.New()
	return uuid.String()
}
