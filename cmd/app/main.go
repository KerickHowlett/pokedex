package main

import (
	c "github.com/KerickHowlett/pokedexcli/cmd/command"
	r "github.com/KerickHowlett/pokedexcli/cmd/repl"
	t "github.com/KerickHowlett/pokedexcli/cmd/toolchain"
)

func main() {
	var toolchain = t.NewToolchain(
		t.WithCommand(c.NewMapCommand()),
		t.WithCommand(c.NewMapBCommand()),
		t.WithCommand(c.NewHelpCommand()),
		t.WithCommand(c.NewExitCommand()),
	)

	repl := r.NewREPL(r.WithCommandExecutor(toolchain.RunCommand))

	repl.StartREPL()
}
