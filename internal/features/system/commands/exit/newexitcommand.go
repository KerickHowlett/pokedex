package exitcommands

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
