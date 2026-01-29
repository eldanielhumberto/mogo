package cmd

import (
	"fmt"

	settingsHelper "github.com/eldanielhumberto/mogo/internal/helpers/settings"
	"github.com/eldanielhumberto/mogo/internal/models"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create mogo.json file",
	Run: func(cmd *cobra.Command, args []string) {
		if settingsHelper.CheckSettingsFileExists() {
			fmt.Printf("The mogo.json already exists")
			return
		}

		settings := &models.Settings{
			Workspaces: map[string]models.Workspace{},
		}

		if err := settingsHelper.SaveSettingsFile(settings); err != nil {
			fmt.Printf("Error saving settings: %v\n", err)
			return
		}

		fmt.Println("init called")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
