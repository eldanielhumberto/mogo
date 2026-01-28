package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mogo",
	Short: "A CLI application for creating monorepos ",
	Long:  `Mogo is a CLI application for easily managing monorepos.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
