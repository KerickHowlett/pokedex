package repl

// StartREPL starts the Read-Eval-Print Loop (REPL) for the REPL struct.
// It continuously prompts the user for input, sanitizes the input, executes the
// command, and prints the result. The loop continues until the user exits the
// REPL.
//
// Example usage:
//
//	repl := NewREPL(WithCommandExecutor(commandExecutor))
//	repl.StartREPL()
//
// The StartREPL method is used to start the REPL.
func (r *REPL) StartREPL() {
	r.PrintPrompt()

	for r.Scanner.Scan() {
		userInput := r.Scanner.Text()
		commands := r.parseInput(userInput)

		r.PrintNewLine()
		r.CommandExecutor(commands[0])

		r.PrintNewLine()
		r.PrintPrompt()
	}
}
