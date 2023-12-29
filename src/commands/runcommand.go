package commands

import "fmt"

// Executes the requested command.
//
// If the requested command is not found, it returns an error.
func RunCommand(command string) error {
	if len(command) == 0 {
		return nil
	}

	selectedCommand, commandFound := GetCommands()[command]

	if !commandFound {
		return fmt.Errorf("Command not found: %s", command)
	}

	return printPaddingWhenRunning(selectedCommand)
}

func printPaddingWhenRunning(command Command) error {
	fmt.Println() // Empty Line
	response := command.Execute()
	fmt.Println() // Empty Line

	return response
}
