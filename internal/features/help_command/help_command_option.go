package help_command

type HelpCommandOption func(*HelpCommand)

// WithDescription is a function that sets the description of the HelpCommand.
//
// The HelpCommandOption function modifies the HelpCommand by setting the
// description to the given description.
//
// Parameters:
//   - description: The description to set for the HelpCommand.
//
// Returns:
//   - An HelpCommandOption function that modifies the HelpCommand by setting
//     the description to the given description.
//
// Example usage:
//
//	command := NewHelpCommand(
//		WithDescription("Exit the Pokemon CLI application."),
//	)
func WithDescription(description string) HelpCommandOption {
	return func(e *HelpCommand) {
		e.description = description
	}
}

// WithName is a function that sets the name of the HelpCommand.
//
// The HelpCommandOption function modifies the HelpCommand by setting the name
// to the given name.
//
// Parameters:
//   - name: The name to set for the HelpCommand.
//
// Returns:
//   - An HelpCommandOption function that modifies the HelpCommand by setting
//     the name to the given name.
//
// Example usage:
//
//	command := NewHelpCommand(
//		WithName("exit"),
//	)
func WithName(name string) HelpCommandOption {
	return func(e *HelpCommand) {
		e.name = name
	}
}
