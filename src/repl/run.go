package repl

import (
	"bufio"
	"fmt"
	"os"
)

// @TODO: Make this function pure.
// @TODO: Make `execute` a more generic type.

// Starts a Read-Eval-Print Loop (REPL) for executing commands.
//
// It takes a function 'execute' as a parameter, which is responsible for
// executing the user input.
func Run(execute func(string) error) {
	const PROMPT = "pokedex > "
	fmt.Print(PROMPT)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		userInput := SanitizeInput(scanner.Text())[0]

		fmt.Println() // Empty Line
		execute(userInput)

		fmt.Println() // Empty Line
		fmt.Print(PROMPT)
	}
}
