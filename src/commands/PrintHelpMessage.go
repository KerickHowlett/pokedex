package commands

import (
	"fmt"
)

// PrintHelpMessage prints the help message for all commands.
// It iterates over the list of commands obtained from GetCommands()
// and calls printCommandHelp() to print the help message for each command.
// It returns an error if there is any issue while printing the help message.
func PrintHelpMessage() error {
	for _, command := range *GetCommands() {
		printCommandHelp(&command)
	}

	return nil
}

func printCommandHelp(command *Command) {
	fmt.Printf("%s: %s\n", command.Name, command.Description)
}
