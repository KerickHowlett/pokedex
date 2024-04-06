package exitcommands

import (
	c "command"
	"os"
)

type ExitCommand struct{}

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
func (e ExitCommand) Execute() error {
	os.Exit(OK)
	return nil
}

func (e ExitCommand) GetDescription() string {
	return "ExitCommand out of the Pokemon CLI application."
}

func (e ExitCommand) GetName() string {
	return "exit"
}

func (e *ExitCommand) PrintHelp() {
	c.PrintHelp(e)
}
