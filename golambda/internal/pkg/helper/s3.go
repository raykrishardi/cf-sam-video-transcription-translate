package helper

import "strings"

func Split(str, delimiter string) string {
	result := str

	parts := strings.Split(str, delimiter)
	if len(parts) > 1 {
		parts = parts[:len(parts)-1]
	}

	result = strings.Join(parts, delimiter)

	return result
}
