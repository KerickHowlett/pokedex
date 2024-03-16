package command

type HelpCommand struct{}

func (c *HelpCommand) Execute() error {
	panic("[ERROR] 'HelpCommand.Execute()' should never be called.")
}

func (c *HelpCommand) GetDescription() string {
	return "Show help for the Pokemon CLI commands."
}

func (c *HelpCommand) GetName() string {
	return "help"
}

func (c *HelpCommand) PrintHelp() {
	printSingleCommandHelp(c)
}

func NewHelpCommand() *HelpCommand {
	return &HelpCommand{}
}
