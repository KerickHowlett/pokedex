package commands

import "fmt"

func RunCommand(command string) error {
	selectedCommand, commandFound := GetCommands()[command]

	if !commandFound {
		return fmt.Errorf("Command not found: %s", command)
	}

	return selectedCommand.Execute()
}
