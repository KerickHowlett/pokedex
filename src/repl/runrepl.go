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

	cli := bufio.NewScanner(os.Stdin)
	for cli.Scan() {
		userInput := parseInput(cli)
		execute(userInput)
		PrintPrompt()
	}
}

func parseInput(cli *bufio.Scanner) string {
	userInput := cli.Text()

	return SanitizeInput(userInput)
}
