package commands

import "fmt"

// Prints the help message for the Pokedex CLI application.
//
// It displays the welcome message and usage instructions, and it prints the
// help information for each command.
func (t Toolchain) PrintHelpMessage() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println() // Empty Line

	fmt.Println("Pokedex Commands")
	fmt.Println("----------------")
	fmt.Println() // Empty Line

	for _, command := range *t.commands {
		command.PrintHelp()
	}
}
