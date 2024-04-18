package help_command

import c "command"

// HelpCommand is a struct that defines the 'help' command for the CLI.
//
// Properties:
//   - name: The name of the command.
//   - description: A brief description of the command.
//
// Methods:
//   - GetName: Returns the name of the HelpCommand.
//   - GetDescription: Returns the description of the HelpCommand.
//   - Execute: Terminates the program with an exit status of 0.
//   - GetArgs: Returns the arguments of the HelpCommand.
//   - SetArgs: Sets the arguments of the HelpCommand.
//   - PrintHelp: Prints the help message for the HelpCommand.
type HelpCommand struct {
	// The name of the command.
	name string
	// A brief description of the command.
	description string
}

// Execute is a method of the Help struct that is responsible for
// executing the actual 'help' command.
//
// It should never be called and will always panic with an error
// message.
//
// Returns:
//   - An error if the method is called.
func (h HelpCommand) Execute() error {
	panic("this method should never be called")
}

// GetArgs returns the arguments of the HelpCommand.
//
// This method is not used by the HelpCommand.
// Its just here to appease the Command static
// type requirements.
//
// Returns:
//   - The arguments of the HelpCommand.
func (h HelpCommand) GetArgs() []string {
	return []string{}
}

// GetDescription returns the description of the HelpCommand.
//
// Returns:
//   - The description of the HelpCommand.
//
// Example usage:
//
//	command := NewHelpCommand()
//	description := command.GetDescription()
func (h HelpCommand) GetDescription() string {
	return h.description
}

// GetName returns the name of the HelpCommand.
//
// Returns:
//   - The name of the HelpCommand.
//
// Example usage:
//
//	command := NewHelpCommand()
//	name := command.GetName()
func (h HelpCommand) GetName() string {
	return h.name
}

// PrintHelp prints the help message for the HelpCommand.
//
// Example usage:
//
//	command := NewHelpCommand()
//	command.PrintHelp()
func (h *HelpCommand) PrintHelp() {
	c.PrintHelp(h)
}

// SetArgs sets the arguments of the HelpCommand.
//
// This method is not used by the HelpCommand.
// Its just here to appease the Command static
// type requirements.
func (h *HelpCommand) SetArgs(args []string) {}
