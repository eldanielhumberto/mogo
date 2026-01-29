package files

import (
	"os"
	"strings"
)

func ParseDirectoryPath(directory string) string {
	runes := []rune(directory)
	if runes[0] != '.' && runes[1] != '/' {
		directory = "./" + directory
	}

	if runes[len(runes)-1] == '/' {
		directory = string(runes[:len(runes)-1])
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
