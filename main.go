package main

import (
	"github.com/KerickHowlett/pokedexcli/src/commands"
	"github.com/KerickHowlett/pokedexcli/src/repl"
)

func main() {
	repl.RunREPL(commands.RunCommand)
}
