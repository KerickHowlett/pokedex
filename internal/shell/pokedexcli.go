package shell

import (
	l "maps/location"
	q "query/state"
	r "repl"
	c "shell/commands"
	t "toochain"
)

func PokedexCLI(config PokedexCLIConfig) {
	mapsState := q.NewQueryState(q.WithNextURL[l.Location](&config.StartingMapsAPIEndpoint))

	toolchain := t.NewToolchain(
		t.WithCommand(c.NewMap(mapsState)),
		t.WithCommand(c.NewMapB(mapsState)),
		t.WithCommand(c.NewHelp()),
		t.WithCommand(c.NewExit()),
	)

	repl := r.NewREPL(
		r.WithCommandExecutor(toolchain.RunCommand),
		r.WithPrompt(config.Prompt),
	)

	repl.RunREPL()
}
