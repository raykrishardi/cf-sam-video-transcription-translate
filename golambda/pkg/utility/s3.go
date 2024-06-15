package utility

import "strings"

func Split(str, delimiter string, getDirPath, getFileName bool) string {
	result := str

	if !getDirPath && !getFileName {
		return result
	}

	parts := strings.Split(str, delimiter)
	if len(parts) > 1 {
		if getDirPath {
			parts = parts[:len(parts)-1]
			result = strings.Join(parts, delimiter)
		} else if getFileName {
			result = parts[len(parts)-1]
		}
	}

	return result
}
