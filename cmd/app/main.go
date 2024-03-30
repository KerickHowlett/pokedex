package main

import (
	c "github.com/KerickHowlett/pokedexcli/cmd/command"
	m "github.com/KerickHowlett/pokedexcli/cmd/mapslist"
	r "github.com/KerickHowlett/pokedexcli/cmd/repl"
	t "github.com/KerickHowlett/pokedexcli/cmd/toolchain"
)

func main() {
	// @TODO: Have API endpoint domain and 'offset' and 'limit' query parameters
	//        be configurable.
	nextURL := "https://pokeapi.co/api/v2/location?offset=0&limit=20"
	mapsState := m.NewMapsList(m.WithNextURL(&nextURL))

	toolchain := t.NewToolchain(
		t.WithCommand(c.NewMapCommand(mapsState)),
		t.WithCommand(c.NewMapBCommand(mapsState)),
		t.WithCommand(c.NewHelpCommand()),
		t.WithCommand(c.NewExitCommand()),
	)

	repl := r.NewREPL(r.WithCommandExecutor(toolchain.RunCommand))
	repl.StartREPL()
}
