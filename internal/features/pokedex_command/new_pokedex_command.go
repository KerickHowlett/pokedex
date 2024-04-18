package pokedex_command

import pc "bills_pc"

func NewPokedexCommand(options ...PokedexCommandOption) *PokedexCommand {
	command := &PokedexCommand{
		description:            "See all the Pokemon you have caught.",
		name:                   "pokedex",
		noCaughtPokemonMessage: "You have not caught any Pokemon yet.",
		pc:                     pc.NewBillsPC(),
		pokedexTitle:           "Your Pokedex:",
	}
	for _, option := range options {
		option(command)
	}
	return command
}
