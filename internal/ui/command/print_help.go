package command

import "fmt"

// CommandDescription represents an interface for getting a brief description of commands.
//
// Methods:
//   - GetDescription: Returns a brief description of the command.
//   - GetName: Returns the name of the command.
type CommandInformation interface {
	CommandDescription
	CommandName
}

// PrintHelp prints the help message for a single command.
//
// Parameters:
//   - command: The command to print the help message for.
//
// Example usage:
//
//	command := ci.NewExitCommand()
//	ci.PrintHelp(command)
func PrintHelp(ci CommandInformation) {
	fmt.Printf("%s: %s\n", ci.GetName(), ci.GetDescription())
}
