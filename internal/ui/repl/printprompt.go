package repl

import "fmt"

const defaultPrompt = "> "

// printPrompt prints the prompt for
// the pokedex REPL.
func (r REPL) printPrompt() {
	fmt.Print(r.prompt)
}
