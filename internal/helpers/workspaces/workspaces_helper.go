package workspaces

import (
	"errors"

	"github.com/eldanielhumberto/mogo/internal/helpers/files"
	settingsHelper "github.com/eldanielhumberto/mogo/internal/helpers/settings"
	"github.com/eldanielhumberto/mogo/internal/models"
)

func AddWorkspace(directory string) error {
	if !files.IsDirectory(directory) {
		return errors.New("Directory " + directory + " is not a valid workspace")
	}

	directoryName := files.ParseDirectoryName(directory)
	settings, err := settingsHelper.ReadSettingsFile()
	if err != nil {
		return err
	}

	if _, ok := settings.Workspaces[directoryName]; ok {
		return err
	}

	newWorkspace := &models.Workspace{
		Context:  directory,
		Commands: make(map[string]string),
	}

	settings.Workspaces[directoryName] = *newWorkspace
	if err := settingsHelper.SaveSettingsFile(settings); err != nil {
		return err
	}

	return nil
}
