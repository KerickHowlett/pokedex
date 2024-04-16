package exit_command

type ExitCommandOption func(*ExitCommand)

// WithName is a function that sets the name of the ExitCommand.
//
// The ExitCommandOption function modifies the ExitCommand by setting the name
// to the given name.
//
// Parameters:
//   - name: The name to set for the ExitCommand.
//
// Returns:
//   - An ExitCommandOption function that modifies the ExitCommand by setting
//     the name to the given name.
//
// Example usage:
//
//	command := NewExitCommand(
//		WithName("exit"),
//	)
func WithName(name string) ExitCommandOption {
	return func(e *ExitCommand) {
		e.name = name
	}
}

// WithDescription is a function that sets the description of the ExitCommand.
//
// The ExitCommandOption function modifies the ExitCommand by setting the
// description to the given description.
//
// Parameters:
//   - description: The description to set for the ExitCommand.
//
// Returns:
//   - An ExitCommandOption function that modifies the ExitCommand by setting
//     the description to the given description.
//
// Example usage:
//
//	command := NewExitCommand(
//		WithDescription("Exit the Pokemon CLI application."),
//	)
func WithDescription(description string) ExitCommandOption {
	return func(e *ExitCommand) {
		e.description = description
	}
}
