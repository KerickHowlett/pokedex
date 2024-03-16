package main

import (
	c "github.com/KerickHowlett/pokedexcli/src/command"
	r "github.com/KerickHowlett/pokedexcli/src/repl"
	t "github.com/KerickHowlett/pokedexcli/src/toolchain"
)

func main() {
	var toolchain = t.NewToolchain(
		t.WithCommand(c.NewExitCommand()),
		t.WithCommand(c.NewHelpCommand()),
	)

	var repl = r.NewREPL(r.WithCommandExecutor(toolchain.RunCommand))

	repl.StartREPL()
}
