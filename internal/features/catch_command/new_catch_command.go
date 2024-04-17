package catch_command

import (
	"math/rand"

	bpc "catch_command/bills_pc"
	p "pokemon"
	"query_fetch"
	"query_fetch/query_cache/ttl"
)

func NewCatchCommand(options ...CatchCommandOption) *CatchCommand {
	command := &CatchCommand{
		cacheTTL:                    ttl.OneDay,
		catchPokemon:                query_fetch.QueryFetch[p.Pokemon],
		checkYourLuck:               func() int { return rand.Int() },
		description:                 "Catch That Pokemon! Usage: catch <pokemon-name>",
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
