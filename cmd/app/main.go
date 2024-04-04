package main

import (
	q "query/state"

	c "github.com/KerickHowlett/pokedexcli/cmd/command"
	l "github.com/KerickHowlett/pokedexcli/cmd/location"
	r "github.com/KerickHowlett/pokedexcli/cmd/repl"
	t "github.com/KerickHowlett/pokedexcli/cmd/toolchain"
)

func main() {
	// @TODO: Have API endpoint domain and 'offset' and 'limit' query parameters
	//        be configurable.
	nextURL := "https://pokeapi.co/api/v2/location?offset=0&limit=20"
	mapsState := q.NewQueryState(
		q.WithNextURL[l.Location](&nextURL),
	)

	toolchain := t.NewToolchain(
		t.WithCommand(c.NewMapCommand(mapsState)),
		t.WithCommand(c.NewMapBCommand(mapsState)),
		t.WithCommand(c.NewHelpCommand()),
		t.WithCommand(c.NewExitCommand()),
	)

	repl := r.NewREPL(
		r.WithCommandExecutor(toolchain.RunCommand),
		r.WithPrompt("pokedex > "),
	)

	repl.RunREPL()
}
