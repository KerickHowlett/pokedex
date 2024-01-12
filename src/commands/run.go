package commands

import "fmt"

// Executes the command selected from the mapped list provided.
//
// If the selection is empty, it does nothing but return nil.
//
// If the requested command is not found, it returns an error.
func Run(selection string, commands Commands) error {
	if len(selection) == 0 {
		return nil
	}

	if selectedCommand, commandFound := commands[selection]; commandFound {
		return selectedCommand.Execute()
	}

	return fmt.Errorf("Command not found: %s", selection)
}
