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
func (r *REPL) StartREPL() {
	r.printPrompt()

	for r.scanner.Scan() {
		userInput := r.scanner.Text()
		commands := r.parseInput(userInput)

		r.printNewLine()
		r.execute(commands[0])

		r.printNewLine()
		r.printPrompt()
	}
}
