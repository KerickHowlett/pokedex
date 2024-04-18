package pokedex_command

import bpc "bills_pc"

type PokedexCommandOption func(*PokedexCommand)

// WithDescription is a function that sets the description of the PokedexCommand.
//
// The PokedexCommandOption function modifies the PokedexCommand by setting the
// description to the given description.
//
// Parameters:
//   - description: The description to set for the PokedexCommand.
//
// Returns:
//   - An PokedexCommandOption function that modifies the PokedexCommand by setting
//     the description to the given description.
//
// Example usage:
//
//	command := NewPokedexCommand(
//		WithDescription("Exit the Pokemon CLI application."),
//	)
func WithDescription(description string) PokedexCommandOption {
	return func(e *PokedexCommand) {
		e.description = description
	}
}

// WithName is a function that sets the name of the PokedexCommand.
//
// The PokedexCommandOption function modifies the PokedexCommand by setting the name
// to the given name.
//
// Parameters:
//   - name: The name to set for the PokedexCommand.
//
// Returns:
//   - An PokedexCommandOption function that modifies the PokedexCommand by setting
//     the name to the given name.
//
// Example usage:
//
//	command := NewHelpCommand(
//		WithName("exit"),
//	)
func WithName(name string) PokedexCommandOption {
	return func(e *PokedexCommand) {
		e.name = name
	}
}

// WithNoCaughtPokemonMessage sets the noCaughtPokemonMessage field in the PokedexCommand struct.
//
// Parameters:
//   - noCaughtPokemonMessage: The message to display when the user has not caught any Pokemon.
//
// Returns:
//   - PokedexCommandOption: A function that sets the noCaughtPokemonMessage field in the PokedexCommand struct.
//
// Example usage:
//
//	c := catch_command.NewPokedexCommand(
//	  catch_command.WithNoCaughtPokemonMessage("You have not caught any Pokemon yet."),
//	)
func WithNoCaughtPokemonMessage(noCaughtPokemonMessage string) PokedexCommandOption {
	return func(c *PokedexCommand) {
		c.noCaughtPokemonMessage = noCaughtPokemonMessage
	}
}

// WithPC sets the pc field in the PokedexCommand struct.
//
// Parameters:
//   - pc: All the Pokemon the user has caught.
//
// Returns:
//   - PokedexCommand: A function that sets the pc field in the PokedexCommand struct.
//
// Example usage:
//
//	c := catch_command.NewPokedexCommand(
//	  catch_command.WithPC(&bpc.BillsPC{}),
//	)
func WithPC(pc *bpc.BillsPC) PokedexCommandOption {
	return func(c *PokedexCommand) {
		c.pc = pc
	}
}

// WithPokedexTitle sets the pokedexTitle field in the PokedexCommand struct.
//
// Parameters:
//   - pokedexTitle: The title of the Pokedex.
//
// Returns:
//   - PokedexCommandOption: A function that sets the pokedexTitle field in the PokedexCommand struct.
//
// Example usage:
//
//	c := catch_command.NewPokedexCommand(
//	  catch_command.WithPokedexTitle("Pokedex"),
//	)
func WithPokedexTitle(pokedexTitle string) PokedexCommandOption {
	return func(c *PokedexCommand) {
		c.pokedexTitle = pokedexTitle
	}
}
