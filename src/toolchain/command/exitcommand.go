package command

import "os"

type ExitCommand struct{}

// Execute terminates the program with an exit status of 0.
//
// Returns:
//   - An error if the program cannot be terminated.
//
// Example usage:
//
//	command := c.NewExitCommand()
//	command.Execute()
func (c *ExitCommand) Execute() error {
	const OKAY_STATUS = 0
	os.Exit(OKAY_STATUS)

	return nil
}

func (c *ExitCommand) GetDescription() string {
	return "Exit out of the Pokemon CLI application."
}

func (c *ExitCommand) GetName() string {
	return "exit"
}

func (c *ExitCommand) PrintHelp() {
	printSingleCommandHelp(c)
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
