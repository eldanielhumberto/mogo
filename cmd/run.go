package cmd

import (
	"fmt"

	"github.com/eldanielhumberto/mogo/internal/helpers/commands"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [workspace] [command]",
	Short: "Execute a command from a workspace",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Usage: mogo run [workspace] [command]")
			return
		}

		workspace := args[0]
		command := args[1]

		if err := commands.RunCommand(workspace, command); err != nil {
			fmt.Printf("Error running command: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
