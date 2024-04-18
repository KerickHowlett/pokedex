package inspect_command

import bpc "bills_pc"

// InspectCommandOption is a type that defines the signature for an option
// function that can be used to configure an InspectCommand.
//
// Parameters:
//   - InspectCommandOption: A function that accepts sets a configurable field
//     for the InspectCommand before instantiating it.
//
// Example usage:
//
//	command := NewInspectCommand()
func NewInspectCommand(options ...InspectCommandOption) *InspectCommand {
	command := &InspectCommand{
		description:                 "Take a closer look at one of your Pokemon. Usage: inspect <pokemon-name>",
		invalidArgsErrorMessage:     "you need to enter a pokemon name to inspect",
		name:                        "inspect",
		noCaughtPokemonMessage:      "You have not caught any Pokemon yet.",
		pc:                          bpc.NewBillsPC(),
		pokemonNotFoundErrorMessage: "you have yet to catch",
	}
	for _, option := range options {
		option(command)
	}
	return command
}
