package help_command

import c "command"

// HelpCommand represents the 'help' command in the CLI.
type HelpCommand struct{}

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
func (m HelpCommand) GetArgs() []string {
	return []string{}
}

// GetDescription returns the description of the HelpCommand.
func (h HelpCommand) GetDescription() string {
	return "Show help for the Pokemon CLI commands."
}

// GetName returns the name of the HelpCommand.
func (h HelpCommand) GetName() string {
	return "help"
}

// PrintHelp prints the help message for the HelpCommand.
func (h *HelpCommand) PrintHelp() {
	c.PrintHelp(h)
}

// SetArgs sets the arguments of the HelpCommand.
func (m *HelpCommand) SetArgs(args []string) {}
