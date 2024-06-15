package utility

import (
	"strings"

	"github.com/google/uuid"
)

func RandomString(length int) string {
	return uuid.NewString()[:length]
}

func GetFileNameOrExtension(fileName string, getExtension bool) string {
	result := fileName

	parts := strings.Split(fileName, ".")
	if len(parts) > 1 {
		if getExtension {
			result = parts[1]
		} else {
			result = parts[0]
		}
	}

	return result
}
