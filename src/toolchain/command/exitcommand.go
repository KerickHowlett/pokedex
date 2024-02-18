package command

import "os"

type ExitCommand struct{}

func (c *ExitCommand) Execute() error {
	const OKAY_STATUS = 0
	os.Exit(OKAY_STATUS)

	return nil
}

func (c *ExitCommand) GetDescription() string {
	return "Exit out of the Pokemon CLI application."
}

func (c *ExitCommand) GetName() string {
	return "exit"
}

func (c *ExitCommand) PrintHelp() {
	printSingleCommandHelp(c)
}
