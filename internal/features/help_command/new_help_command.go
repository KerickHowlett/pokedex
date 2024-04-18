package help_command

// NewHelpCommand creates a new instance of the Help struct.
//
// Parameters:
//   - options: A list of HelpCommandOption functions to apply to the Help struct.
//
// Returns:
//   - A new instance of the Help struct.
//
// Example usage:
//
//	command := NewHelpCommand()
func NewHelpCommand(options ...HelpCommandOption) *HelpCommand {
	command := &HelpCommand{
		description: "Show help for the Pokemon CLI commands.",
		name:        "help",
	}
	for _, option := range options {
		option(command)
	}
	return command
}
