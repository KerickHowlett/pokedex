package repl

import "fmt"

// RunREPL starts the Read-Eval-Print Loop (REPL) for the REPL struct.
// It continuously prompts the user for input, sanitizes the input, executes the
// command, and prints the result. The loop continues until the user exits the
// REPL.
//
// Example usage:
//
//	repl := NewREPL(WithCommandExecutor(commandExecutor))
//	repl.RunREPL()
func (r *REPL) RunREPL() {
	r.printPrompt()

	for r.scanner.Scan() {
		r.printNewLine()

		userInput := r.parseInput(r.scanner.Text())
		command, args := userInput[0], userInput[1:]
		if err := r.execute(command, args); err != nil {
			fmt.Println(err)
			fmt.Println("Please try again.")
		}

		r.printNewLine()
		r.printPrompt()
	}
}
