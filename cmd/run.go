package cmd

import (
	"fmt"

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
		fmt.Printf("Executing command '%s' in workspace '%s'\n", command, workspace)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
