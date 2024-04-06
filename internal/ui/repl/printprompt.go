package repl

import "fmt"

// defaultPrompt is the default prompt string
// used in the REPL (Read-Eval-Print Loop).
const defaultPrompt = "> "

// printPrompt prints the REPL prompt to the console.
//
// Example usage:
//
//	repl := NewREPL()
//	repl.printPrompt()
func (r REPL) printPrompt() {
	fmt.Print(r.prompt)
}
