package main

import (
	"github.com/KerickHowlett/pokedexcli/commands"
	"github.com/KerickHowlett/pokedexcli/repl"
)

func main() {
	repl.RunREPL(commands.RunCommand)
}
