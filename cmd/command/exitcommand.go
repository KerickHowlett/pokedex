package command

import (
	c "command"
	"os"
)

type ExitCommand struct{}

const okayStatus = 0

// Execute terminates the program with an exit status of 0.
//
// Returns:
//   - An error if the program cannot be terminated.
//
// Example usage:
//
//	command := c.NewExitCommand()
//	command.Execute()
func (e ExitCommand) Execute() error {
	os.Exit(okayStatus)
	return nil
}

func (e ExitCommand) GetDescription() string {
	return "Exit out of the Pokemon CLI application."
}

func (e ExitCommand) GetName() string {
	return "exit"
}

func (e *ExitCommand) PrintHelp() {
	c.PrintHelp(e)
}

// NewExitCommand creates a new instance of ExitCommand.
// When the 'exit' command is executed, it will exit the application.
//
// Returns:
//   - A new instance of the ExitCommand struct.
//
// Example usage:
//
//	command := NewExitCommand()
func NewExitCommand() *ExitCommand {
	return &ExitCommand{}
}
