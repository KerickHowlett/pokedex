package pokedex_command

import (
	pc "bills_pc"
	c "command"
	"fmt"
)

// PokedexCommand is the command to view the list of caught Pokemon.
//
// Fields:
//   - description: The description of the PokedexCommand.
//   - name: The unique identifier for the PokedexCommand.
//   - noCaughtPokemonMessage: The message to display when the user has not caught any Pokemon.
//   - pc: The state that stores the user's caught Pokemon.
//   - pokedexTitle: The title for the Pokedex.
//
// Methods:
//   - Execute: Executes the PokedexCommand.
//   - GetArgs: Returns the arguments of the PokedexCommand.
//   - GetDescription: Returns the description of the PokedexCommand.
//   - GetName: Returns the name of the PokedexCommand.
//   - PrintHelp: Prints the help message for the PokedexCommand.
//   - SetArgs: Sets the arguments of the PokedexCommand.
type PokedexCommand struct {
	// The description of the PokedexCommand.
	description string
	// The unique identifier for the PokedexCommand.
	name string
	// The message to display when the user has not caught any Pokemon.
	noCaughtPokemonMessage string
	// The BillsPC stores the user's caught Pokemon.
	pc *pc.BillsPC
	// The title that displays above the list of caught Pokemon.
	pokedexTitle string
}

// Execute is a method of the Help struct that is responsible for
// executing the actual 'pokedex' command.
//
// It should never be called and will always panic with an error
// message.
//
// Returns:
//   - An error if the method is called.
func (p PokedexCommand) Execute() error {
	if p.pc.TotalCaughtPokemon() == 0 {
		fmt.Println(p.noCaughtPokemonMessage)
		return nil
	}

	fmt.Println(p.pokedexTitle)
	for pokemonName := range p.pc.GetCaughtPokemon() {
		fmt.Printf("- %s\n", pokemonName)
	}

	return nil
}

// GetArgs returns the arguments of the PokedexCommand.
//
// This method is not used by the PokedexCommand.
// Its just here to appease the Command static
// type requirements.
//
// Returns:
//   - The arguments of the PokedexCommand.
func (p PokedexCommand) GetArgs() []string {
	return []string{}
}

// GetDescription returns the description of the PokedexCommand.
//
// Returns:
//   - The description of the PokedexCommand.
//
// Example usage:
//
//	command := NewPokedexCommand()
//	description := command.GetDescription()
func (p PokedexCommand) GetDescription() string {
	return p.description
}

// GetName returns the name of the PokedexCommand.
//
// Returns:
//   - The name of the PokedexCommand.
//
// Example usage:
//
//	command := NewPokedexCommand()
//	name := command.GetName()
func (p PokedexCommand) GetName() string {
	return p.name
}

// PrintHelp prints the help message for the PokedexCommand.
//
// Example usage:
//
//	command := NewPokedexCommand()
//	command.PrintHelp()
func (p *PokedexCommand) PrintHelp() {
	c.PrintHelp(p)
}

// SetArgs sets the arguments of the PokedexCommand.
//
// This method is not used by the PokedexCommand.
// Its just here to appease the Command static
// type requirements.
func (p PokedexCommand) SetArgs(args []string) {}
