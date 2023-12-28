package repl

import (
	"bufio"
	"os"

	"github.com/KerickHowlett/pokedexcli/src/commands"
)

func RunREPL() {
	PrintPrompt()

	cli := bufio.NewScanner(os.Stdin)
	for cli.Scan() {
		userInput := parseInput(cli)
		commands.RunCommand(userInput)

		PrintPrompt()
	}
}

func parseInput(cli *bufio.Scanner) string {
	userInput := cli.Text()

	return SanitizeInput(userInput)
}
