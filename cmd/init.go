package cmd

import (
	"fmt"

	settingsHelper "github.com/eldanielhumberto/mogo/internal/helpers/settings"
	"github.com/eldanielhumberto/mogo/internal/models"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create mogo.json file",
	Run: func(cmd *cobra.Command, args []string) {
		if settingsHelper.CheckSettingsFileExists() {
			pterm.Println("The mogo.json already exists")
			return
		}

		pterm.Println()
		pterm.DefaultBigText.WithLetters(putils.LettersFromString("MOGO REPO")).Render()
		pterm.Println()
		pterm.Println()

		settings := &models.Settings{
			Workspaces: map[string]models.Workspace{},
		}

		if err := settingsHelper.SaveSettingsFile(settings); err != nil {
			fmt.Printf("Error saving settings: %v\n", err)
			return
		}

		pterm.Println("'mogo.json' created successfully")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
