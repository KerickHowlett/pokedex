package exit_command

import (
	"os"

	c "command"
)

// ExitCommand is a struct that defines the 'exit' command for the CLI.
//
// Properties:
//   - name: The name of the command.
//   - description: A brief description of the command.
//
// Methods:
//   - GetName: Returns the name of the ExitCommand.
//   - GetDescription: Returns the description of the ExitCommand.
//   - Execute: Terminates the program with an exit status of 0.
//   - GetArgs: Returns the arguments of the ExitCommand.
//   - SetArgs: Sets the arguments of the ExitCommand.
//   - PrintHelp: Prints the help message for the ExitCommand.
type ExitCommand struct {
	// The name of the command.
	name string
	// A brief description of the command.
	description string
}

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

// GetArgs returns the arguments of the ExitCommand.
func (e ExitCommand) GetArgs() []string {
	return []string{}
}

// GetDescription returns the description of the ExitCommand.
func (e ExitCommand) GetDescription() string {
	return e.description
}

// GetName returns the name of the ExitCommand.
func (e ExitCommand) GetName() string {
	return e.name
}

// PrintHelp prints the help message for the ExitCommand.
func (e *ExitCommand) PrintHelp() {
	c.PrintHelp(e)
}

// SetArgs sets the arguments of the ExitCommand.
func (e *ExitCommand) SetArgs(args []string) {}
