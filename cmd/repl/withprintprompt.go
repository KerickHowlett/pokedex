package repl

// WithPrintPrompt sets the function to be called for printing the REPL prompt.
// It is an option function that can be used with the REPL struct.
//
// Parameters:
//   - printPrompt: The function to be called for printing the REPL prompt.
//
// Returns:
//   - An option function that sets the printPrompt function for the REPL.
//
// Example usage:
//
//	printPrompt := func() {
//		fmt.Print("pokedex > ")
//	}
//	repl := NewREPL(WithPrintPrompt(printPrompt))
//
// Now the printPrompt function will be used to print the REPL prompt.
func WithPrintPrompt(printPrompt func()) REPLOption {
	return func(r *REPL) {
		r.PrintPrompt = printPrompt
	}
}
