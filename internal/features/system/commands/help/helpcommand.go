package helpcommands

import c "command"

type Help struct{}

// Execute is a method of the Help struct that is responsible for
// executing the actual 'help' command.
//
// It should never be called and will always panic with an error message.
func (h Help) Execute() error {
	panic("[ERROR] 'Help.Execute()' should never be called.")
}

func (h Help) GetDescription() string {
	return "Show help for the Pokemon CLI commands."
}

func (h Help) GetName() string {
	return "help"
}

func (h *Help) PrintHelp() {
	c.PrintHelp(h)
}
