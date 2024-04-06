package command

// Command represents a command that can be executed in the CLI.
type Command interface {
	// Execute executes the command and returns an error if any.
	Execute() error

	// GetDescription returns a brief description of the command.
	GetDescription() string

	// GetName returns the name of the command.
	GetName() string

	// PrintHelp prints the help information for the command.
	PrintHelp()
}
