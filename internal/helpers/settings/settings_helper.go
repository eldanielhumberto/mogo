package settings

import (
	"encoding/json"
	"os"

	"github.com/eldanielhumberto/mogo/internal/constants"
	"github.com/eldanielhumberto/mogo/internal/helpers/files"
	"github.com/eldanielhumberto/mogo/internal/models"
)

func CheckSettingsFileExists() bool {
	return files.CheckFileExists(constants.SETTINGS_FILE)
}

func ReadSettingsFile() (*models.Settings, error) {
	file, err := os.ReadFile(constants.SETTINGS_FILE)
	if err != nil {
		return nil, err
	}

	settings := &models.Settings{}
	json.Unmarshal(file, &settings)

	return settings, nil
}

func SaveSettingsFile(settings *models.Settings) error {
	jsonData, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(constants.SETTINGS_FILE, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
