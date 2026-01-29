package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/eldanielhumberto/mogo/internal/helpers"
	"github.com/spf13/cobra"
)

type Settings struct {
	Commands map[string]any `json:"commands"`
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create mogo.json file",
	Run: func(cmd *cobra.Command, args []string) {
		filename := "mogo.json"
		exists := helpers.CheckFileExists(filename)

		if exists {
			fmt.Printf("File %s already exists\n", filename)
			return
		}

		settings := Settings{
			Commands: map[string]any{},
		}
		jsonData, err := json.MarshalIndent(settings, "", "  ")
		if err != nil {
			fmt.Printf("Error marshalling to JSON: %v\n", err)
			return
		}

		err = os.WriteFile("mogo.json", jsonData, 0644)
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
