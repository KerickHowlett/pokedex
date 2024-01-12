package main

import (
	"github.com/KerickHowlett/pokedexcli/commands"
	"github.com/KerickHowlett/pokedexcli/repl"
)

func main() {
	cmds := commands.Create()
	commandRunner := func(selection string) error {
		return commands.Run(selection, cmds)
	}

	repl.Run(commandRunner)
}
