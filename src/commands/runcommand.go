package commands

import "fmt"

// Executes the requested command.
//
// If the requested command is not found, it returns an error.
func RunCommand(command string) error {
	if len(command) == 0 {
		return nil
	}

	if selectedCommand, commandFound := GetCommands()[command]; commandFound {
		return selectedCommand.Execute()
	}

	return fmt.Errorf("Command not found: %s", command)
}
