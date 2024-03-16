package toolchain

import "fmt"

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
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println() // Add new line.

	fmt.Println("Pokedex Commands")
	fmt.Println("--------------------")
	fmt.Println() // Add new line.

	// Print help info for all toolchain commands.
	for _, command := range *t.commands {
		(*command).PrintHelp()
	}
}
