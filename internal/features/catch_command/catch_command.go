package catch_command

import (
	"fmt"
	qec "query_fetch/query_cache/cache_eviction_config"

	bpc "bills_pc"
	cmd "command"
	p "pokemon"
	qf "query_fetch"
)

// CatchPokemonFunc is a type alias for the QueryFetchFunc type that fetches Pokemon from the Pokemon API.
type CatchPokemonFunc qf.QueryFetchFunc[p.Pokemon]

// CheckYourLuckFunc is a type alias for a function that returns the chance of catching a Pokemon.
type CheckYourLuckFunc func(pokemonBaseExperience int) (chance int)

// CatchCommand represents a command related to catching Pokemon in the pokedexcli application.
//
// Fields:
//   - args: The arguments provided to the command.
//   - apiEndpoint: The URL of the API endpoint to fetch the location area data.
//   - catchPokemon: A function that fetches locations/maps from the Pokemon API.
//   - checkYourLuck: A function that returns the chance of catching a Pokemon.
//   - description: Describes the purpose of the ExploreCommand.
//   - difficulty: Effects the challenge level at catching any given Pokemon.
//   - ec: The eviction configuration for the query cache.
//   - name: The name of the ExploreCommand.
//   - noEnteredArgsErrorMessage: An error message to display when no arguments are entered.
//   - noEncountersFoundErrorMessage: An error message to display when no maps are found.
//   - pc: Records all the Pokemon the user has caught.
//   - successfulCatchNotification: A message to display when a Pokemon is caught.
//   - throwBallNotification: A message to display when a Pokemon is thrown a ball.
//
// Methods:
//   - Execute: Attempts to catch a Pokemon.
//   - GetArgs: Returns the arguments provided to the CatchCommand.
//   - GetDescription: Returns the description of the CatchCommand.
//   - GetName: Returns the name of the CatchCommand.
//   - PrintHelp: Prints the help message for the CatchCommand.
//   - SetArgs: Sets the arguments for the CatchCommand.
//   - isCatchSuccessful: Checks if the CatchCommand is successful in catching a Pokemon.
//   - hasValidArgs: Checks if the CatchCommand has valid arguments needed to run the Execute() method properly.
type CatchCommand struct {
	// The arguments provided to the command.
	args []string
	// The URL of the API endpoint to fetch the location area data.
	apiEndpoint string
	// catchPokemon is a function that fetches locations/maps from the Pokemon API.
	catchPokemon CatchPokemonFunc
	// checkYourLuck is a function that returns the chance of catching a Pokemon.
	checkYourLuck CheckYourLuckFunc
	// description describes the purpose of the ExploreCommand.
	description string
	// difficulty effects the challenge level at catching any given Pokemon.
	difficulty int
	// ec is the eviction configuration for the query cache.
	ec *qec.QueryEvictionConfig
	// escapedNotification is a message to display when a Pokemon escapes.
	escapedNotification string
	// name is the name of the ExploreCommand.
	name string
	// noEncountersFoundErrorMessage is an error message to display when no maps are found.
	noEnteredArgsErrorMessage string
	// pc records all the Pokemon the user has caught.
	pc *bpc.BillsPC
	// successfulCatchNotification is a message to display when a Pokemon is caught.
	successfulCatchNotification string
	// throwBallNotification is a message to display when a Pokemon is thrown a ball.
	throwBallNotification string
}

// CatchCommand.Execute attempts to catch a Pokemon.
//
// Returns:
//   - An error if the command fails to execute.
//
// Example usage:
//
//	command := NewCatchCommand()
//	err := command.Execute()
func (c *CatchCommand) Execute() error {
	if !c.hasValidArgs() {
		return fmt.Errorf(c.noEnteredArgsErrorMessage)
	}
	pokemonName := c.args[0]

	// Threw Pokeball
	fmt.Printf("%s %s...\n", c.throwBallNotification, pokemonName)
	wildPokemon, err := c.catchPokemon(c.apiEndpoint+"/"+pokemonName, c.ec)
	if err != nil {
		return err
	}

	if !c.isCatchSuccessful(wildPokemon) {
		// Pokemon Escaped
		fmt.Printf("%s %s\n", pokemonName, c.escapedNotification)
		return nil
	}

	// Caught Pokemon
	fmt.Printf("%s %s!\n", c.successfulCatchNotification, pokemonName)
	c.pc.Deposit(wildPokemon)

	return nil
}

// GetArgs returns the arguments provided to the CatchCommand.
//
// Returns:
//   - The arguments provided to the CatchCommand.
//
// Example usage:
//
//	command := NewCatchCommand()
//	args := command.GetArgs()
func (c *CatchCommand) GetArgs() []string {
	return c.args
}

// GetDescription returns the description of the CatchCommand.
//
// Returns:
//   - The description of the CatchCommand.
//
// Example usage:
//
//	command := NewCatchCommand()
//	description := command.GetDescription()
func (c *CatchCommand) GetDescription() string {
	return c.description
}

// GetName returns the name of the CatchCommand.
//
// Returns:
//   - The name of the CatchCommand.
//
// Example usage:
//
//	command := NewCatchCommand()
//	name := command.GetName()
func (c *CatchCommand) GetName() string {
	return c.name
}

// PrintHelp prints the help message for the CatchCommand.
//
// Example usage:
//
//	command := NewCatchCommand()
//	command.PrintHelp()
func (c *CatchCommand) PrintHelp() {
	cmd.PrintHelp(c)
}

// SetArgs sets the arguments for the CatchCommand.
//
// Parameters:
//   - args: The arguments provided to the command.
//
// Example usage:
//
//	command := NewCatchCommand()
//	command.SetArgs([]string{"1"})
func (c *CatchCommand) SetArgs(args []string) {
	c.args = args
}

// isCatchSuccessful checks if the CatchCommand is successful in catching a Pokemon.
//
// Returns:
//   - A boolean indicating if the CatchCommand is successful in catching a Pokemon.
//
// Example usage:
//
//	command := NewCatchCommand()
//	isSuccessful := command.isCatchSuccessful()
func (c *CatchCommand) isCatchSuccessful(wildPokemon *p.Pokemon) bool {
	return c.checkYourLuck(wildPokemon.BaseExperience) > c.difficulty
}

// hasValidArgs checks if the CatchCommand has valid arguments needed to run the Execute() method properly.
//
// Returns:
//   - A boolean indicating if the CatchCommand has valid arguments.
//
// Example usage:
//
//	command := NewCatchCommand()
//	command.SetArgs([]string{"1"})
//	hasValidArgs := command.hasValidArgs()
func (c CatchCommand) hasValidArgs() bool {
	return c.args != nil && len(c.args) > 0 && c.args[0] != ""
}
