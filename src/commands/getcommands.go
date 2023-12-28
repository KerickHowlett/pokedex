package commands

type Commands = map[string]Command

// GetCommands returns a map of commands available in the CLI.
// Each command is represented by a key-value pair, where the key is the command
// name and the value is a Command struct containing the command's name,
// description, and execute function.
func GetCommands() Commands {
	return Commands{
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
