package catch_command

import (
	"math/rand"

	bpc "bills_pc"
	p "pokemon"
	qf "query_fetch"
	ec "query_fetch/query_cache/cache_eviction_config"
)

// NewCatchCommand returns a new instance of the CatchCommand struct.
//
// Parameters:
//   - options: The options to set on the CatchCommand.
//
// Returns:
//   - A new instance of the CatchCommand struct.
//
// Example usage:
//
//	command := NewCatchCommand()
func NewCatchCommand(options ...CatchCommandOption) *CatchCommand {
	command := &CatchCommand{
		catchPokemon:                qf.QueryFetch[p.Pokemon],
		checkYourLuck:               defaultCheckYourLuck,
		description:                 "Catch That Pokemon! Usage: catch <pokemon-name>",
		difficulty:                  50,
		ec:                          ec.NewQueryEvictionConfig(),
		escapedNotification:         "escaped!",
		name:                        "catch",
		noEnteredArgsErrorMessage:   "a pokemon name is required",
		pc:                          bpc.NewBillsPC(),
		successfulCatchNotification: "You caught",
		throwBallNotification:       "You threw a pokeball at",
	}

	for _, option := range options {
		option(command)
	}

	if command.apiEndpoint == "" {
		panic("an api endpoint is required")
	}

	return command
}

// defaultCheckYourLuck is the default implementation of the CheckYourLuck function.
//
// Parameters:
//   - pokemonBaseExperience: The base experience of the pokemon to catch.
//
// Returns:
//   - The chance of catching the Pokemon.
//
// Example usage:
//
//	chance := defaultCheckYourLuck(100)
func defaultCheckYourLuck(pokemonBaseExperience int) (chance int) {
	chance = rand.Intn(pokemonBaseExperience)
	return chance
}
