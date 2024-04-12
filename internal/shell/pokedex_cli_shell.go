package pokedexclishell

import (
	e "explore/commands/explore_command"
	mf "maps/commands/map"
	mb "maps/commands/mapb"
	ms "maps/state"
	r "repl"
	ex "system/commands/exit_command"
	h "system/commands/help_command"
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
	helpCommand := h.NewHelpCommand()
	exitCommand := ex.NewExitCommand()

	sharedMapState := ms.NewMapsState(ms.WithNextURL(&config.StartingMapsAPIEndpoint))
	mapCommand := mf.NewMapCommand(sharedMapState)
	mapBCommand := mb.NewMapBCommand(sharedMapState)

	exploreCommand := e.NewExploreCommand(e.WithAPIEndpoint(config.LocalAreaAPIEndpoint))

	toolchain := t.NewToolchain(
		t.WithCommand(exploreCommand),
		t.WithCommand(mapCommand),
		t.WithCommand(mapBCommand),
		t.WithCommand(helpCommand),
		t.WithCommand(exitCommand),
	)

	repl := r.NewREPL(
		r.WithCommandExecutor(toolchain.RunCommand),
		r.WithPrompt(config.Prompt),
	)

	repl.RunREPL()
}
