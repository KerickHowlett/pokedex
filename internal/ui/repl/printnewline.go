package repl

import "fmt"

// printNewLine prints a new line to the console.
//
// Example usage:
//
//	repl := NewREPL()
//	repl.printNewLine()
func (r REPL) printNewLine() {
	fmt.Println()
}
