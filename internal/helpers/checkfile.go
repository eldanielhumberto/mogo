package helpers

import (
	"os"

	"github.com/eldanielhumberto/mogo/internal/constants"
)

func CheckFileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}

	return true
}

func CheckSettingsFileExists() bool {
	return CheckFileExists(constants.SETTINGS_FILE)
}
