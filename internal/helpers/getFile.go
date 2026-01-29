package helpers

import (
	"encoding/json"
	"os"

	"github.com/eldanielhumberto/mogo/internal/constants"
	"github.com/eldanielhumberto/mogo/internal/models"
)

func GetSettingsFile() (*models.Settings, error) {
	file, err := os.ReadFile(constants.SETTINGS_FILE)
	if err != nil {
		return nil, err
	}

	settings := models.Settings{}
	if err := json.Unmarshal(file, &settings); err != nil {
		return nil, err
	}

	return &settings, nil
}
