package repl

// REPL represents a Read-Eval-Print Loop for the command-line interface.
//
// Fields:
//   - execute: A function that takes a string as input and returns an error.
//   - prompt: A string representing the text displayed to the user as a command prompt.
//   - scanner: A scanner interface responsible for reading user input from the command line.
type REPL struct {
	// execute is a function that takes a string as input and returns an error.
	// It is responsible for executing the given command string.
	//
	// Parameters:
	//   - selection: The command string to execute.
	//	 - args: A slice of strings representing the arguments to the command.
	//
	// Returns:
	//   - An error if any.
	execute func(selection string, args []string) error
	// prompt represents the text that is displayed to the user as a command prompt.
	// It is used to prompt the user for input in the REPL (Read-Eval-Print Loop).
	prompt string
	// scanner is responsible for reading user input from the command line.
	scanner scanner
}
