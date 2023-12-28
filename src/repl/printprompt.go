package repl

import "fmt"

// Prints the prompt for the PokedexCLI REPL.
func PrintPrompt() {
	const prompt = "pokedex > "

	fmt.Print(prompt)
}
