package inspect_command

import (
	"fmt"

	bpc "bills_pc"
	c "command"
)

// InspectCommand is the command used a Pokemon the user has caught previously.
//
// Properties:
//   - args: The arguments provided to the command.
//   - description: Explains what the command does and how to use it.
//   - name: Is the command's unique identifier.
//   - pc: Bill's PC that stores all the Pokemon the user has caught.
//
// Methods:
//   - GetName: Returns the name of the InspectCommand.
//   - GetDescription: Returns the description of the InspectCommand.
//   - Execute: Terminates the program with an exit status of 0.
//   - GetArgs: Returns the arguments of the InspectCommand.
//   - SetArgs: Sets the arguments of the InspectCommand.
//   - PrintHelp: Prints the help message for the InspectCommand.
type InspectCommand struct {
	// args are the arguments provided to the command.
	args []string
	// description explains what the command does and how to use it.
	description string
	// name is this command's unique identifier.
	name string
	// invalidArgsErrorMessage is the message displayed when the user provides invalid arguments.
	invalidArgsErrorMessage string
	// noCaughtPokemonMessage is the message displayed when the user has not caught any Pokemon.
	noCaughtPokemonMessage string
	// pc records all the Pokemon the user has caught.
	pc *bpc.BillsPC
	// pokemonNotFoundErrorMessage is the message displayed when the user tries to inspect a Pokemon they haven't caught.
	pokemonNotFoundErrorMessage string
}

// Execute inspects a Pokemon caught by the user that's selected by it's name.
//
// Returns:
//   - An error if the method is called.
//
// Example usage:
//
//	command := NewInspectCommand()
//	err := command.Execute()
func (i InspectCommand) Execute() error {
	if i.pc.TotalCaughtPokemon() == 0 {
		fmt.Println(i.noCaughtPokemonMessage)
		return nil
	}

	if !i.hasValidArgs() {
		return fmt.Errorf(i.invalidArgsErrorMessage)
	}

	pokemonName := i.args[0]
	pokemon, isPokemonCaught := i.pc.Inspect(pokemonName)
	if !isPokemonCaught {
		return fmt.Errorf("%s %s", i.pokemonNotFoundErrorMessage, pokemonName)
	}

	printBasicInfoOf(pokemon)
	printStatsOf(pokemon)
	printTypesOf(pokemon)

	return nil
}

// GetArgs returns the arguments of the InspectCommand.
//
// Returns:
//   - The arguments of the InspectCommand.
//
// Example usage:
//
//	command := NewInspectCommand()
//	args := command.GetArgs()
func (i InspectCommand) GetArgs() []string {
	return i.args
}

// GetDescription returns the description of the InspectCommand.
//
// Returns:
//   - The description of the InspectCommand.
//
// Example usage:
//
//	command := NewInspectCommand()
//	description := command.GetDescription()
func (i InspectCommand) GetDescription() string {
	return i.description
}

// GetName returns the name of the InspectCommand.
//
// Returns:
//   - The name of the InspectCommand.
//
// Example usage:
//
//	command := NewInspectCommand()
//	name := command.GetName()
func (i InspectCommand) GetName() string {
	return i.name
}

// PrintHelp prints the help message for the InspectCommand.
//
// Example usage:
//
//	command := NewInspectCommand()
//	command.PrintHelp()
func (i *InspectCommand) PrintHelp() {
	c.PrintHelp(i)
}

// SetArgs sets the arguments of the InspectCommand.
//
// Parameters:
//   - args: The arguments provided to the command.
//
// Example usage:
//
//	command := NewInspectCommand()
//	command.SetArgs([]string{"pikachu"})
func (i *InspectCommand) SetArgs(args []string) {
	i.args = args
}

// hasValidArgs checks if the InspectCommand has valid arguments needed to run the Execute() method properly.
//
// Returns:
//   - A boolean indicating if the InspectCommand has valid arguments.
//
// Example usage:
//
//	command := NewInspectCommand()
//	command.SetArgs([]string{"1"})
//	hasValidArgs := command.hasValidArgs()
func (i InspectCommand) hasValidArgs() bool {
	return i.args != nil && len(i.args) > 0 && i.args[0] != ""
}
