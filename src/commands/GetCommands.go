package commands

func GetCommands() *map[string]Command {
	return &map[string]Command{
		"exit": {
			Name:        "exit",
			Description: "Exits the REPL",
			Execute:     ExitREPL,
		},
		"help": {
			Name:        "help",
			Description: "Prints this help message",
			Execute:     PrintHelpMessage,
		},
	}
}
