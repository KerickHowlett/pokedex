package toolchain

import "fmt"

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
