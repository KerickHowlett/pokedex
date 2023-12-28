package commands

import "fmt"

func RunCommand(command string) error {
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
