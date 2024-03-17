package main

import (
	c "github.com/KerickHowlett/pokedexcli/internal/command"
	r "github.com/KerickHowlett/pokedexcli/internal/repl"
	t "github.com/KerickHowlett/pokedexcli/internal/toolchain"
)

func main() {
	var toolchain = t.NewToolchain(
		t.WithCommand(c.NewExitCommand()),
		t.WithCommand(c.NewHelpCommand()),
	)

	var repl = r.NewREPL(r.WithCommandExecutor(toolchain.RunCommand))

	repl.StartREPL()
}
