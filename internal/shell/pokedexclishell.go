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

// PokedexCLIShell is a function that initializes and runs the Pokedex CLI shell.
// It takes a PokedexCLIConfig as input and sets up the necessary components for
// the shell.
// The shell provides commands to interact with the Pokedex, such as displaying
// maps, executing map commands, displaying help information, and exiting the
// shell.
//
// Parameters:
//   - config: A PokedexCLIConfig struct that contains the configuration options for the Pokedex CLI.
//
// Example usage:
//
//	config := PokedexCLIConfig{
//	  StartingMapsAPIEndpoint: "https://api.example.com/maps",
//	  Prompt: "> ",
//	}
//	PokedexCLIShell(config)
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