package utils

import (
	"github.com/google/uuid"
)

func GetRandomString(length int) string {
	return uuid.NewString()[:length]
}
