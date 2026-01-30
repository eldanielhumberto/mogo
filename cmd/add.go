package cmd

import (
	"fmt"

	"github.com/eldanielhumberto/mogo/internal/helpers/files"
	"github.com/eldanielhumberto/mogo/internal/helpers/settings"
	"github.com/eldanielhumberto/mogo/internal/helpers/workspaces"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [workspace_folder]",
	Short: "Add a workspace",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 || args[0] == "" || args[0] == "." {
			fmt.Printf("Usage: mogo add [workspace_folder]\n\n")
			return
		}

		if !settings.CheckSettingsFileExists() {
			fmt.Println("Settings file not found")
			return
		}

		directory := files.ParseDirectoryPath(args[0])
		if err := workspaces.AddWorkspace(directory); err != nil {
			fmt.Printf("Error adding workspace: %s\n\n", err)
			return
		}

		fmt.Printf("Adding workspace at %s\n\n", directory)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
