package pokedexclishell

import (
	e "explore"
	m "maps/commands"
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
	exitCommand := ex.NewExitCommand()
	helpCommand := h.NewHelpCommand()

	sharedMapState := ms.NewMapsState(ms.WithNextURL(&config.StartingMapsAPIEndpoint))
	mapCommand := m.NewMapsCommand(
		m.WithPaginationDirection(m.Next),
		m.WithState(sharedMapState),
	)
	mapBCommand := m.NewMapsCommand(
		m.WithPaginationDirection(m.Previous),
		m.WithState(sharedMapState),
	)

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
