package repl

import (
	"fmt"
	"testing"

	m "github.com/KerickHowlett/pokedexcli/tests/mocks"
	"github.com/KerickHowlett/pokedexcli/tests/utils"
)

func TestStartREPL(t *testing.T) {
	fmt.Println("StartREPL should execute command and prompt for input until user exits")
	executedCommand, expectedCommand := setupStartREPLTest()

	utils.ExpectSameEntity(t, *executedCommand, expectedCommand, "executedCommand")
}

func setupStartREPLTest() (executedCommand *string, userInput string) {
	executedCommand = new(string)
	scanner := m.NewScannerMock()

	mockedCommandExecutor := func(args string) error {
		*executedCommand = args
		defer func() { scanner.SetIsEnabled(false) }()
		return nil
	}

	repl := NewREPL(
		WithCommandExecutor(mockedCommandExecutor),
		WithParseInput(m.ParserMock),
		WithPrintNewLine(emptyFunctionMock),
		WithPrintPrompt(emptyFunctionMock),
		WithScanner(scanner),
	)

	repl.StartREPL()
	userInput = scanner.Text()

	return
}
