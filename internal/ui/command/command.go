package command

// Command represents a command that can be executed in the CLI.
//
// @TODO: Break up into smaller interfaces.
//
// Methods:
//   - Execute: Executes the command and returns an error if any.
//   - GetArgs: Returns the arguments passed to the command.
//   - GetDescription: Returns a brief description of the command.
//   - GetName: Returns the name of the command.
//   - PrintHelp: Prints the help information for the command.
//   - SetArgs: Sets the arguments for the command.
type Command interface {
	// Execute executes the command and returns an error if any.
	// Returns:
	//   - An error if any.
	Execute() error
	// GetArgs returns the arguments passed to the command.
	// Returns:
	//   - A slice of strings representing the command's entered arguments.
	GetArgs() []string
	// GetDescription returns a brief description of the command.
	// Returns:
	//   - A string representing the description of the command.
	GetDescription() string
	// GetName returns the name of the command.
	// Returns:
	//   - A string representing the name of the command.
	GetName() string
	// PrintHelp prints the help information for the command.
	// Returns:
	//   - A string representing the help information.
	PrintHelp()
	// SetArgs sets the arguments for the command.
	// Parameters:
	//   - args: A slice of strings representing the arguments.
	SetArgs(args []string)
}

// Commands is a map that stores command names as keys
// and their corresponding Command structs as values.
type Commands = map[string]*Command
