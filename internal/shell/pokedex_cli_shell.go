package pokedexclishell

import (
	catch "catch_command"
	bpc "catch_command/bills_pc"
	exit "exit_command"
	explore "explore_command"
	help "help_command"
	mc "map_command"
	pd "map_command/pagination_direction"
	ms "map_command/state"
	"repl"
	tc "toochain"
)

// @TODO: Add means to shutdown CLI gracefully with a signal handler.

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
	exitCommand := exit.NewExitCommand()
	helpCommand := help.NewHelpCommand()

	sharedMapState := ms.NewMapsState(ms.WithNextURL(&config.StartingMapsAPIEndpoint))
	mapCommand := mc.NewMapCommand(
		mc.WithPaginationDirection(pd.Next),
		mc.WithState(sharedMapState),
	)
	mapBCommand := mc.NewMapCommand(
		mc.WithPaginationDirection(pd.Previous),
		mc.WithState(sharedMapState),
	)

	exploreCommand := explore.NewExploreCommand(
		explore.WithAPIEndpoint(config.LocalAreaAPIEndpoint),
	)

	billsPC := bpc.NewBillsPC()
	catchCommand := catch.NewCatchCommand(
		catch.WithAPIEndpoint(config.PokemonAPIEndpoint),
		catch.WithPC(billsPC),
	)

	toolchain := tc.NewToolchain(
		tc.WithCommand(catchCommand),
		tc.WithCommand(exploreCommand),
		tc.WithCommand(mapCommand),
		tc.WithCommand(mapBCommand),
		tc.WithCommand(helpCommand),
		tc.WithCommand(exitCommand),
	)

	repl := repl.NewREPL(
		repl.WithCommandExecutor(toolchain.RunCommand),
		repl.WithPrompt(config.Prompt),
	)

	repl.RunREPL()
}
