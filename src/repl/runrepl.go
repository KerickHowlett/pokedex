package repl

import (
	"bufio"
	"os"
)

type CommandDelegate = func(...string) error

func RunREPL(execute CommandDelegate) {
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
