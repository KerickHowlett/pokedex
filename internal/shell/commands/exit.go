package commands

import (
	c "command"
	"os"
)

type Exit struct{}

const OK = 0

// Execute terminates the program with an exit status of 0.
//
// Returns:
//   - An error if the program cannot be terminated.
//
// Example usage:
//
//	command := c.NewExit()
//	command.Execute()
func (e Exit) Execute() error {
	os.Exit(OK)
	return nil
}

func (e Exit) GetDescription() string {
	return "Exit out of the Pokemon CLI application."
}

func (e Exit) GetName() string {
	return "exit"
}

func (e *Exit) PrintHelp() {
	c.PrintHelp(e)
}

// NewExit creates a new instance of Exit.
// When the 'exit' command is executed, it will exit the application.
//
// Returns:
//   - A new instance of the Exit struct.
//
// Example usage:
//
//	command := NewExit()
func NewExit() *Exit {
	return &Exit{}
}
