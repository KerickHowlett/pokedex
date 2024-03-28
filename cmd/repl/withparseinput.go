package repl

// WithParseInput is a REPLOption function that sets the ParseInput
// function for the REPL.
//
// The ParseInput function takes a string as input and returns a slice of
// strings after sanitizing the input.
//
// It is used to sanitize the user input before processing it in the REPL.
//
// Parameters:
//   - ParseInput: The ParseInput function to be set for the REPL.
//
// Returns:
//   - An option function that sets the ParseInput function for the REPL.
//
// Example usage:
//
//	ParseInput := func(input string) []string {
//		return strings.Fields(input)
//	}
//	repl := NewREPL(WithParseInput(ParseInput))
//
// Now the ParseInput function will be used to sanitize user input during the
// REPL session.
func WithParseInput(ParseInput func(string) []string) REPLOption {
	return func(r *REPL) {
		r.ParseInput = ParseInput
	}
}
