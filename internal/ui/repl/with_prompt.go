package repl

// WithPrompt sets the prompt for the REPL.
//
// Parameters:
//   - prompt: The string to be displayed as the prompt in the REPL.
//
// Returns:
//   - An option function that sets the prompt for the REPL.
//
// Example usage:
//
//	repl := NewREPL(
//		WithPrompt("Enter command: "),
//	)
//
// Now the prompt "Enter command: " will be displayed to the user in the REPL.
func WithPrompt(prompt string) REPLOption {
	return func(r *REPL) {
		r.prompt = prompt
	}
}
