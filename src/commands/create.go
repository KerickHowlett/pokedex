package commands

type Commands = map[string]Command

// A map of available commands is returned.
//
// Each command is represented by a key-value pair, where the key is the command
// name and the value is a Command struct containing the command's name,
// description, and function to be executed.
func Create() Commands {
	cmds := Commands{
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

	// @NOTE: This is a minor hack to get around the fact Go doesn't allow cmds
	//        to "pass itself" as an argument to `PrintHelpMessage` before being
	//        declared.
	helpCommand := cmds["help"]
	helpCommand.Execute = func() error {
		return PrintHelpMessage(cmds)
	}

	cmds["help"] = helpCommand

	return cmds
}
