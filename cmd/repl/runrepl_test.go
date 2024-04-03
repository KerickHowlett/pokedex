package repl

import (
	"testing"

	s "testtools/mocks/scanner"
)

func TestStartREPL(t *testing.T) {
	setup := func() (executedCommand *string, userInput string) {
		executedCommand = new(string)
		scanner := s.NewMockScanner()
		scanner.SetUserInput("mock")

		mockedCommandExecutor := func(args string) error {
			*executedCommand = args
			defer scanner.SetIsEnabled(false)
			return nil
		}

		repl := NewREPL(
			WithCommandExecutor(mockedCommandExecutor),
			WithScanner(scanner),
		)

		repl.RunREPL()
		userInput = scanner.Text()

		return executedCommand, userInput
	}

	t.Run("should execute command and prompt for input until user exits", func(t *testing.T) {
		if executedCommand, expectedCommand := setup(); *executedCommand != expectedCommand {
			t.Errorf("Executed command does not match expected command. Got: %s, Expected: %s", *executedCommand, expectedCommand)
		}
	})
}
