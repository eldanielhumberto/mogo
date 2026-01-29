package files

import (
	"os"
	"path/filepath"
	"strings"
)

func ParseDirectoryPath(directory string) string {
	directory = filepath.Clean(directory)

	if filepath.IsAbs(directory) || directory == "." {
		return directory
	}

	prefix := "." + string(filepath.Separator)
	if !strings.HasPrefix(directory, prefix) {
		directory = prefix + directory
	}

	return directory
}

func ParseDirectoryName(directory string) string {
	return strings.Split(directory, "/")[len(strings.Split(directory, "/"))-1]
}

func CheckFileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}

	return true
}

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
