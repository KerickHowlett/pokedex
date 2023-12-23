package commands

import (
	"fmt"
)

func PrintHelpMessage() error {
	for _, command := range *GetCommands() {
		printCommandHelp(&command)
	}

	return nil
}

func printCommandHelp(command *Command) {
	fmt.Printf("%s: %s\n", command.Name, command.Description)
}
