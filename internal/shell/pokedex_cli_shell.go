package pokedexclishell

import (
	exc "exit"
	epc "explore"
	mc "map"
	ms "map/state"
	"repl"
	h "system/commands/help_command"
	toolchain "toochain"
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
	exitCommand := exc.NewExitCommand()
	helpCommand := h.NewHelpCommand()

	sharedMapState := ms.NewMapsState(ms.WithNextURL(&config.StartingMapsAPIEndpoint))
	mapCommand := mc.NewMapCommand(
		mc.WithPaginationDirection(mc.Next),
		mc.WithState(sharedMapState),
	)
	mapBCommand := mc.NewMapCommand(
		mc.WithPaginationDirection(mc.Previous),
		mc.WithState(sharedMapState),
	)

	exploreCommand := epc.NewExploreCommand(epc.WithAPIEndpoint(config.LocalAreaAPIEndpoint))

	toolchain := toolchain.NewToolchain(
		toolchain.WithCommand(exploreCommand),
		toolchain.WithCommand(mapCommand),
		toolchain.WithCommand(mapBCommand),
		toolchain.WithCommand(helpCommand),
		toolchain.WithCommand(exitCommand),
	)

	repl := repl.NewREPL(
		repl.WithCommandExecutor(toolchain.RunCommand),
		repl.WithPrompt(config.Prompt),
	)

	repl.RunREPL()
}
