package repl

import "fmt"

// defaultPrintPrompt is the
// default PrintPrompt function
// for the REPL.
func defaultPrintPrompt() {
	fmt.Print("pokedex > ")
}
