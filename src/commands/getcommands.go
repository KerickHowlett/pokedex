package commands

type Commands = map[string]Command

// A map of available commands is returned.
//
// Each command is represented by a key-value pair, where the key is the command
// name and the value is a Command struct containing the command's name,
// description, and function to be executed.
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
