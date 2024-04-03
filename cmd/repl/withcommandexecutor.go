package repl

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
func WithCommandExecutor(commandExecutor func(string) error) REPLOption {
	return func(r *REPL) {
		r.execute = commandExecutor
	}
}
