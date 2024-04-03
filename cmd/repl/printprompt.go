package repl

import "fmt"

// printPrompt prints the prompt for
// the pokedex REPL.
func (r REPL) printPrompt() {
	fmt.Print("pokedex > ")
}
