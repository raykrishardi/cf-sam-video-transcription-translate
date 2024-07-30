package utils

import "strings"

func GetFileNameOrExtension(file string, getExtension bool) string {
	result := file

	parts := strings.Split(file, ".")
	if len(parts) > 1 {
		if getExtension {
			result = parts[1]
		} else {
			result = parts[0]
		}
	}

	return result
}

func GetDirPathOrFileName(path, delimiter string, getDirPath, getFileName bool) string {
	result := path

	if !getDirPath && !getFileName {
		return result
	}

	parts := strings.Split(path, delimiter)
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
