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
	CommandArguments
	CommandDescription
	CommandExecutor
	CommandHelp
	CommandName
}

// CommandArguments represents an interface for command arguments.
//
// Methods:
//   - GetArgs: Returns the arguments passed to the command.
//   - SetArgs: Sets the arguments for the command.
type CommandArguments interface {
	// GetArgs returns the arguments passed to the command.
	// Returns:
	//   - A slice of strings representing the command's entered arguments.
	GetArgs() []string
	// SetArgs sets the arguments for the command.
	// Parameters:
	//   - args: A slice of strings representing the arguments.
	SetArgs(args []string)
}

// CommandDescription represents an interface for getting a brief description of commands.
//
// Methods:
//   - GetDescription: Returns a brief description of the command.
type CommandDescription interface {
	// GetDescription returns a brief description of the command.
	// Returns:
	//   - A string representing the description of the command.
	GetDescription() string
}

// CommandExecutor represents an interface for executing commands.
//
// Methods:
//   - Execute: Executes the command and returns an error if any.
type CommandExecutor interface {
	// Execute executes the command and returns an error if any.
	// Returns:
	//   - An error if any.
	Execute() error
}

// CommandHelp represents an interface for printing help information for commands.
//
// Methods:
//   - PrintHelp: Prints the help information for the command.
type CommandHelp interface {
	// PrintHelp prints the help information for the command.
	PrintHelp()
}

// CommandHelp represents an interface for printing help information for commands.
//
// Methods:
//   - PrintHelp: Prints the help information for the command.
type CommandName interface {
	// GetName returns the name of the command.
	// Returns:
	//   - A string representing the name of the command.
	GetName() string
}

// Commands is a map that stores command names as keys
// and their corresponding Command structs as values.
type Commands = map[string]*Command
