package repl

import (
	"bufio"
	"os"
)

// Starts a Read-Eval-Print Loop (REPL) for executing commands.
//
// It takes a function 'execute' as a parameter, which is responsible for
// executing the user input.
func RunREPL(execute func(string) error) {
	PrintPrompt()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		userInput := SanitizeInput(scanner.Text())[0]

		if len(userInput) == 0 {
			continue
		}

		execute(userInput)
		PrintPrompt()
	}
}
