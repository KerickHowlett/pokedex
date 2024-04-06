package command

import "fmt"

// PrintHelp prints the help message for a single command.
//
// Parameters:
//   - command: The command to print the help message for.
//
// Example usage:
//
//	command := c.NewExitCommand()
//	c.PrintHelp(command)
func PrintHelp(command Command) {
	name := command.GetName()
	description := command.GetDescription()

	fmt.Printf("%s: %s\n", name, description)
}
