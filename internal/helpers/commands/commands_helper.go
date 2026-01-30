package commands

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/eldanielhumberto/mogo/internal/helpers/settings"
	"github.com/pterm/pterm"
)

type Task struct {
	Workspace string
	Cmd       *exec.Cmd
}

type prefixWriter struct {
	context string
}

func (pw *prefixWriter) Write(p []byte) (n int, err error) {
	pterm.NewRGB(15, 199, 209).Printf("%s | ", pw.context)
	pterm.DefaultParagraph.Printf("%s ", string(p))
	pterm.Println()

	return len(p), nil
}

func RunCommand(workspace, command string) error {
	pterm.DefaultSection.Printf("Command '%s' in workspace '%s'", command, workspace)
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

func RunCommandInParallel(command string) error {
	settings, err := settings.ReadSettingsFile()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	tasks := make(chan *exec.Cmd)

	for range len(settings.Workspaces) {
		wg.Add(1)
		go worker(tasks, &wg)
	}

	for _, w := range settings.Workspaces {
		if cmdStr, ok := w.Commands[command]; ok {
			pw := &prefixWriter{
				context: w.Context,
			}

			parts := strings.Fields(cmdStr)
			cmd := exec.Command(parts[0], parts[1:]...)
			cmd.Dir = w.Context
			cmd.Stdout = pw
			cmd.Stderr = pw

			tasks <- cmd
		}
	}

	close(tasks)
	wg.Wait()
	return nil
}

func worker(tasks <-chan *exec.Cmd, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		if err := task.Run(); err != nil {
			fmt.Printf("Error executing command '%s': %v\n\n", task.Args[0], err)
		}
	}
}
