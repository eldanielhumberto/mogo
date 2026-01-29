package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/eldanielhumberto/mogo/internal/constants"
	"github.com/eldanielhumberto/mogo/internal/helpers"
	"github.com/eldanielhumberto/mogo/internal/models"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create mogo.json file",
	Run: func(cmd *cobra.Command, args []string) {
		if helpers.CheckSettingsFileExists() {
			fmt.Printf("The mogo.json already exists")
			return
		}

		settings := models.Settings{
			Workspaces: map[string]models.Workspace{},
		}
		jsonData, err := json.MarshalIndent(settings, "", "  ")
		if err != nil {
			fmt.Printf("Error marshalling to JSON: %v\n", err)
			return
		}

		err = os.WriteFile(constants.SETTINGS_FILE, jsonData, 0644)
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			return
		}

		fmt.Println("init called")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
