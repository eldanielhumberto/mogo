package helpers

import "os"

func IsDirectory(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}
