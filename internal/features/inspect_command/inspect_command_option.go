package inspect_command

import bpc "bills_pc"

type InspectCommandOption func(*InspectCommand)

// WithDescription sets the description of the InspectCommand.
//
// Parameters:
//   - description: The description of the InspectCommand.
//
// Returns:
//   - An InspectCommandOption that sets the description of the InspectCommand.
//
// Example usage:
//
//	command := NewInspectCommand(
//		WithDescription("Inspect the bills PC"),
//	)
func WithDescription(description string) InspectCommandOption {
	return func(i *InspectCommand) {
		i.description = description
	}
}

// WithInvalidArgsErrorMessage sets the invalidArgsErrorMessage field in the InspectCommand struct.
//
// Parameters:
//   - invalidArgsErrorMessage: The error message to display when the user enters invalid arguments.
//
// Returns:
//   - InspectCommandOption: A function that sets the invalidArgsErrorMessage field in the InspectCommand struct.
//
// Example usage:
//
//	c := catch_command.NewInspectCommand(
//	  catch_command.WithInvalidArgsErrorMessage("You need to enter a Pokemon name to inspect"),
//	)
func WithInvalidArgsErrorMessage(invalidArgsErrorMessage string) InspectCommandOption {
	return func(c *InspectCommand) {
		c.invalidArgsErrorMessage = invalidArgsErrorMessage
	}
}

// WithName sets the name of the InspectCommand.
//
// Parameters:
//   - name: The name of the InspectCommand.
//
// Returns:
//   - An InspectCommandOption that sets the name of the InspectCommand.
//
// Example usage:
//
//	command := NewInspectCommand(
//		WithName("inspect"),
//	)
func WithName(name string) InspectCommandOption {
	return func(i *InspectCommand) {
		i.name = name
	}
}

// WithNoCaughtPokemonMessage sets the noCaughtPokemonMessage field in the InspectCommand struct.
//
// Parameters:
//   - noCaughtPokemonMessage: The message to display when the user has not caught any Pokemon.
//
// Returns:
//   - InspectCommandOption: A function that sets the noCaughtPokemonMessage field in the InspectCommand struct.
//
// Example usage:
//
//	c := catch_command.NewInspectCommand(
//	  catch_command.WithNoCaughtPokemonMessage("You have not caught any Pokemon yet."),
//	)
func WithNoCaughtPokemonMessage(noCaughtPokemonMessage string) InspectCommandOption {
	return func(c *InspectCommand) {
		c.noCaughtPokemonMessage = noCaughtPokemonMessage
	}
}

// WithPC sets the pc field in the InspectCommand struct.
//
// Parameters:
//   - pc: All the Pokemon the user has caught.
//
// Returns:
//   - InspectCommandOption: A function that sets the pc field in the InspectCommand struct.
//
// Example usage:
//
//	c := catch_command.NewInspectCommand(
//	  catch_command.WithPC(&bpc.BillsPC{}),
//	)
func WithPC(pc *bpc.BillsPC) InspectCommandOption {
	return func(c *InspectCommand) {
		c.pc = pc
	}
}

// WithPokemonNotFoundErrorMessage sets the pokemonNotFoundErrorMessage field in the InspectCommand struct.
//
// Parameters:
//   - pokemonNotFoundErrorMessage: The message to display when the user tries to inspect a Pokemon they haven't caught.
//
// Returns:
//   - InspectCommandOption: A function that sets the pokemonNotFoundErrorMessage field in the InspectCommand struct.
//
// Example usage:
//
//	c := catch_command.NewInspectCommand(
//	  catch_command.WithPokemonNotFoundErrorMessage("You have yet to catch"),
//	)
func WithPokemonNotFoundErrorMessage(pokemonNotFoundErrorMessage string) InspectCommandOption {
	return func(c *InspectCommand) {
		c.pokemonNotFoundErrorMessage = pokemonNotFoundErrorMessage
	}
}
