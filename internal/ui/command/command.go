package command

// Command represents a command that can be executed in the CLI.
//
// Methods:
//   - Execute: Executes the command and returns an error if any.
//   - GetDescription: Returns a brief description of the command.
//   - GetName: Returns the name of the command.
//   - PrintHelp: Prints the help information for the command.
type Command interface {
	Execute() error
	GetDescription() string
	GetName() string
	PrintHelp()
}
