package repl

import (
	"fmt"
	"testing"

	s "testtools/mocks/scanner"
)

func TestStartREPL(t *testing.T) {
	fmt.Println("StartREPL should execute command and prompt for input until user exits")
	executedCommand, expectedCommand := setupStartREPLTest()

	if *executedCommand != expectedCommand {
		t.Errorf("Executed command does not match expected command. Got: %s, Expected: %s", *executedCommand, expectedCommand)
	}
}

func setupStartREPLTest() (executedCommand *string, userInput string) {
	executedCommand = new(string)
	scanner := s.NewMockScanner()

	mockedCommandExecutor := func(args string) error {
		*executedCommand = args
		defer func() { scanner.SetIsEnabled(false) }()
		return nil
	}

	repl := NewREPL(
		WithCommandExecutor(mockedCommandExecutor),
		WithPrintNewLine(emptyFunctionMock),
		WithPrintPrompt(emptyFunctionMock),
		WithScanner(scanner),
	)

	repl.StartREPL()
	userInput = scanner.Text()

	return
}
