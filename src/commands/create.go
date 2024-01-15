package commands

type Commands = map[string]Command

// A map of available commands is returned.
//
// Each command is represented by a key-value pair, where the key is the command
// name and the value is a Command struct containing the command's name,
// description, and function to be executed.
func Create() Commands {
	return Commands{
		"exit": {
			Name:        "exit",
			Description: "Exits the REPL",
			Execute:     ExitProcess,
		},
		"help": {
			Name:        "help",
			Description: "Prints this help message",
			Execute:     func() error { return nil },
		},
	}
}

func NewCommand(name string, description string, execute func() error) *Command {
	return &Command{
		Name:        name,
		Description: description,
		Execute:     execute,
	}
}
