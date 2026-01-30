package commands

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/eldanielhumberto/mogo/internal/helpers/settings"
)

func RunCommand(workspace, command string) error {
	fmt.Printf("Excute command '%s' in workspace '%s'\n\n", command, workspace)

	settings, err := settings.ReadSettingsFile()
	if err != nil {
		return err
	}

	if _, ok := settings.Workspaces[workspace]; !ok {
		return errors.New("Workspace '" + workspace + "' does not exist")
	}

	if _, ok := settings.Workspaces[workspace].Commands[command]; !ok {
		return errors.New("Command '" + command + "' does not exist in workspace '" + workspace + "'")
	}

	contextCommand := settings.Workspaces[workspace].Context
	parseCommand := strings.Fields(settings.Workspaces[workspace].Commands[command])

	cmd := exec.Command(parseCommand[0], parseCommand[1:]...)
	cmd.Dir = contextCommand
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error executing command '%s' in workspace '%s': %v\n\n", command, workspace, err)
		return err
	}

	return nil
}
