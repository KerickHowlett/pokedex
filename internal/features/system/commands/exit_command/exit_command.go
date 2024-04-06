package exit_command

import (
	c "command"
	"os"
)

type ExitCommand struct{}

// Execute terminates the program with an exit status of 0.
//
// Returns:
//   - An error if the program cannot be terminated.
//
// Example usage:
//
//	command := c.NewExit()
//	command.Execute()
func (e ExitCommand) Execute() error {
	os.Exit(OK)
	return nil
}

// GetDescription returns the description of the ExitCommand.
func (e ExitCommand) GetDescription() string {
	return "ExitCommand out of the Pokemon CLI application."
}

// GetName returns the name of the ExitCommand.
func (e ExitCommand) GetName() string {
	return "exit"
}

// PrintHelp prints the help message for the ExitCommand.
func (e *ExitCommand) PrintHelp() {
	c.PrintHelp(e)
}
