package repl

// REPLOption is a function type that represents options for configuring a REPL.
type REPLOption func(*REPL)

// WithCommandExecutor is a function that sets the command executor for the REPL.
//
// The command executor is responsible for executing commands in the REPL.
//
// It takes a variadic argument of strings representing the command and returns
// an error if any.
//
// This function is used as an option in the REPL constructor to customize the
// behavior of the REPL.
//
// Parameters:
//   - commandExecutor: The function to be called for executing commands in the
//     REPL.
//
// Returns:
//   - An option function that sets the commandExecutor function for the REPL.
//
// Example usage:
//
//	commandExecutor := func(args ...string) error {
//		fmt.Println("Executing command:", args)
//		return nil
//	}
//	repl := NewREPL(WithCommandExecutor(commandExecutor))
//
// Now the commandExecutor function will be used to execute commands in the REPL.
func WithCommandExecutor(commandExecutor func(command string, args []string) error) REPLOption {
	return func(r *REPL) {
		r.execute = commandExecutor
	}
}

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
