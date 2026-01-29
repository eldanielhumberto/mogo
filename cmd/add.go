package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/eldanielhumberto/mogo/internal/constants"
	"github.com/eldanielhumberto/mogo/internal/helpers"
	"github.com/eldanielhumberto/mogo/internal/models"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a workspace",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 || args[0] == "" || args[0] == "." {
			fmt.Println("Please provide a directory")
			return
		}

		// Check if settings file exists
		if !helpers.CheckSettingsFileExists() {
			fmt.Println("Settings file not found")
			return
		}

		// Delete trailing slash
		directory := args[0]
		runes := []rune(directory)
		if runes[0] != '.' && runes[1] != '/' {
			directory = "./" + directory
		}

		if runes[len(runes)-1] == '/' {
			directory = string(runes[:len(runes)-1])
		}

		// Check if directory is a valid workspace
		if !helpers.IsDirectory(directory) {
			fmt.Printf("Directory %s is not a valid workspace\n", directory)
			return
		}

		// Check if workspace already exists
		directoryName := strings.Split(directory, "/")[len(strings.Split(directory, "/"))-1]
		file, err := os.ReadFile(constants.SETTINGS_FILE)
		if err != nil {
			fmt.Println("Error in to read file ", err)
			return
		}

		settings := &models.Settings{}
		json.Unmarshal(file, &settings)

		if _, ok := settings.Workspaces[directoryName]; ok {
			fmt.Printf("Workspace %s already exists\n", directory)
			return
		}

		// Add workspace
		newWorkspace := &models.Workspace{
			Context:  directory,
			Commands: make(map[string]string),
		}

		settings.Workspaces[directoryName] = *newWorkspace

		dataBytes, err := json.MarshalIndent(settings, "", "  ")
		if err != nil {
			fmt.Println("Error ", err)
			return
		}

		err = os.WriteFile(constants.SETTINGS_FILE, dataBytes, 0644)
		if err != nil {
			fmt.Println("Error ", err)
			return
		}

		fmt.Printf("Adding workspace at %s\n", directory)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
