package repl

// WithScanner sets the scanner for the REPL.
// It is an option function that can be used when creating a new REPL instance.
// The scanner is responsible for reading user input during the REPL session.
//
// Parameters:
//   - scanner: The scanner to be set for the REPL.
//
// Returns:
//   - An option function that sets the scanner for the REPL.
//
// Example usage:
//
//	scanner := bufio.NewScanner(os.Stdin)
//	repl := NewREPL(WithScanner(scanner))
//
// Now the scanner will be used to read user input during the REPL session.
func WithScanner(scanner scanner) REPLOption {
	return func(r *REPL) {
		r.scanner = scanner
	}
}
