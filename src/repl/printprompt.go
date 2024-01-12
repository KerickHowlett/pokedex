package repl

import "fmt"

// Prints the prompt for the PokedexCLI REPL.
func PrintPrompt() {
	const PROMPT = "pokedex > "
	fmt.Print(PROMPT)
}
