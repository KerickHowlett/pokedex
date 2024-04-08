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
		userInput := r.scanner.Text()
		command := r.parseInput(userInput)[0]

		r.printNewLine()
		if err := r.execute(command); err != nil {
			fmt.Println(err)
			fmt.Println("Please try again.")
		}

		r.printNewLine()
		r.printPrompt()
	}
}
