package main

import (
	l "maps/location"
	q "query/state"

	r "repl"
	t "tc"

	cmd "github.com/KerickHowlett/pokedexcli/cmd/commands"
	conf "github.com/KerickHowlett/pokedexcli/cmd/config"
)

func main() {
	mapsState := q.NewQueryState(q.WithNextURL[l.Location](&conf.StartingMapsAPIEndpoint))

	toolchain := t.NewToolchain(
		t.WithCommand(cmd.NewMapCommand(mapsState)),
		t.WithCommand(cmd.NewMapBCommand(mapsState)),
		t.WithCommand(cmd.NewHelpCommand()),
		t.WithCommand(cmd.NewExitCommand()),
	)

	repl := r.NewREPL(
		r.WithCommandExecutor(toolchain.RunCommand),
		r.WithPrompt("pokedex > "),
	)

	repl.RunREPL()
}
