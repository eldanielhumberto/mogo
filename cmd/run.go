package cmd

import (
	"fmt"

	"github.com/eldanielhumberto/mogo/internal/helpers/commands"
	"github.com/eldanielhumberto/mogo/internal/helpers/settings"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [command]",
	Short: "Execute a command from a workspace",
	Run: func(cmd *cobra.Command, args []string) {
		if !settings.CheckSettingsFileExists() {
			fmt.Println("Settings file not found")
			return
		}

		if len(args) == 0 {
			fmt.Println("Usage: mogo run [command] [flags]")
			return
		}

		command := args[0]
		workspace, _ := cmd.Flags().GetString("workspace")
		if workspace == "" {
			fmt.Printf("Excute command '%s' in parallel\n\n", args[0])
			return
		}

		if err := commands.RunCommand(workspace, command); err != nil {
			fmt.Printf("Error running command: %v\n", err)
		}
	},
}

func init() {
	runCmd.Flags().StringP("workspace", "w", "", "workspace name")
	rootCmd.AddCommand(runCmd)
}
