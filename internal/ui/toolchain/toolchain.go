package toolchain

import (
	"fmt"

	c "command"
	hmp "toochain/help_message_prints"
)

// Toolchain represents a collection of commands used in the UI toolchain.
//
// Fields:
//   - commands: A pointer to a map of Commands.
//   - prints: A pointer to a HelpMessagePrints struct.
type Toolchain struct {
	// commands is a pointer to a map of Commands.
	commands *c.Commands
	// prints is a pointer to a HelpMessagePrints struct.
	prints *hmp.HelpMessagePrints
}

// PrintHelpMessage prints the help message for the Pokedex toolchain.
// It prints the welcome message and the help info for all toolchain CLI commands.
//
// Example usage:
//
//	toolchain := NewToolchain(
//		WithCommand(c.NewExitCommand()),
//		WithCommand(c.NewHelpCommand()),
//	)
//
//	toolchain.PrintHelpMessage()
func (t *Toolchain) PrintHelpMessage() {
	fmt.Println(t.prints.Title)
	t.addEmptyLine()

	fmt.Println(t.prints.CommandsSubTitle)
	fmt.Println(t.prints.HR)
	t.addEmptyLine()

	t.printAllCommandHelps()
}

// RunCommand executes the specified command based on the given selection.
//
// It returns an error if the selection is not valid or if there is an error
// executing the command.
//
// Parameters:
//   - selection: The name of the command to execute.
//   - args: A slice of strings representing the arguments to the command.
//
// Returns:
//   - An error if the selection is not valid or if there is an error executing
//     the command.
//
// Example usage:
//
//	toolchain := NewToolchain(
//		WithCommand(c.NewExitCommand()),
//		WithCommand(c.NewHelpCommand()),
//	)
//
//	err := toolchain.RunCommand("exit")
func (t *Toolchain) RunCommand(selection string, args []string) error {
	if selection == "" {
		return nil
	}

	if selection == "help" {
		t.PrintHelpMessage()
		return nil
	}

	if selectedCommand, commandFound := t.SelectCommand(selection); commandFound {
		(*selectedCommand).SetArgs(args)
		return (*selectedCommand).Execute()
	}

	return fmt.Errorf("command '%s' is not valid", selection)
}

// SelectCommand plucks a CLI command from the toolchain based on the given
// command name.
//
// It returns the found command and a boolean indicating whether the command
// exists.
//
// Parameters:
//   - commandName: The name of the command to select.
//
// Returns:
//   - The found command.
//   - A boolean indicating whether the command exists.
//
// Example usage:
//
//	toolchain := NewToolchain(
//		WithCommand(c.NewExitCommand()),
//		WithCommand(c.NewHelpCommand()),
//	)
func (t *Toolchain) SelectCommand(commandName string) (foundCommand *c.Command, commandExists bool) {
	foundCommand, commandExists = (*t.commands)[commandName]
	return foundCommand, commandExists
}

// addEmptyLine prints an empty line to the console for spacing for better readability.
//
// Example usage:
//
//	toolchain.addEmptyLine()
func (t *Toolchain) addEmptyLine() {
	fmt.Println()
}

// printAllCommandHelps prints the help info for all toolchain CLI commands.
//
// Example usage:
//
//	toolchain := NewToolchain(
//		WithCommand(c.NewExitCommand()),
//		WithCommand(c.NewHelpCommand()),
//	)
//
//	toolchain.printAllCommandHelps()
func (t *Toolchain) printAllCommandHelps() {
	for _, command := range *t.commands {
		(*command).PrintHelp()
	}
}
