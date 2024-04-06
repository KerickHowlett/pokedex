package pokedexclishell

import (
	m "maps/commands/map"
	mb "maps/commands/mapb"
	l "maps/location"
	q "query/state"
	r "repl"
	e "system/commands/exit"
	h "system/commands/help"
	t "toochain"
)

func PokedexCLIShell(config PokedexCLIConfig) {
	mapsState := q.NewQueryState(q.WithNextURL[l.Location](&config.StartingMapsAPIEndpoint))

	toolchain := t.NewToolchain(
		t.WithCommand(m.NewMapCommand(mapsState)),
		t.WithCommand(mb.NewMapBCommand(mapsState)),
		t.WithCommand(h.NewHelpCommand()),
		t.WithCommand(e.NewExitCommand()),
	)

	repl := r.NewREPL(
		r.WithCommandExecutor(toolchain.RunCommand),
		r.WithPrompt(config.Prompt),
	)

	repl.RunREPL()
}
