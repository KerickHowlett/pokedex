package exit_command

// NewExitCommand creates a new instance of ExitCommand.
// When the 'exit' command is executed, it will exit the application.
//
// Returns:
//   - A new instance of the ExitCommand struct.
//
// Example usage:
//
//	command := NewExitCommand()
func NewExitCommand(options ...ExitCommandOption) *ExitCommand {
	command := &ExitCommand{
		name:        "exit",
		description: "Exit the Pokemon CLI application.",
	}
	for _, option := range options {
		option(command)
	}
	return command
}
