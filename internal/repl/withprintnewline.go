package repl

// WithPrintNewLine is a REPLOption function that sets the PrintNewLine function
// of the REPL.
//
// The PrintNewLine function is called to print an empty line in the REPL.
//
// It takes a printEmptyLine function as a parameter and assigns it to the
// PrintNewLine field of the REPL.
//
// Parameters:
//   - printEmptyLine: The function to be called for printing an empty line.
//
// Returns:
//   - An option function that sets the printEmptyLine function for the REPL.
//
// Example usage:
//
//	printEmptyLine := func() {
//		fmt.Println()
//	}
//	repl := NewREPL(WithPrintNewLine(printEmptyLine))
//
// Now the printEmptyLine function will be used to print an empty line in the
// REPL.
func WithPrintNewLine(printEmptyLine func()) REPLOption {
	return func(r *REPL) {
		r.PrintNewLine = printEmptyLine
	}
}
