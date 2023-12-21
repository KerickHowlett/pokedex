package commands

var commands = map[string]CLICommand{
	"help": {
		Name:        "help",
		Description: "Prints this help message",
		Callback:    PrintHelp,
	},
	"exit": {
		Name:        "exit",
		Description: "Exits the program",
		Callback:    ExitRepl,
	},
}

func Commands() {

}
