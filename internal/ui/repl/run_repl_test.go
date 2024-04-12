package repl

import (
	"fmt"
	"strings"
	"testing"

	s "test_tools/mocks/scanner"
	"test_tools/utils"
)

func TestStartREPL(t *testing.T) {
	setup := func(responseType string) (executedCommand *string, userInput string, output string) {
		executedCommand = new(string)

		scanner := s.NewMockScanner()
		mockedCommandExecutor := func(command string, args []string) error {
			*executedCommand = command
			defer scanner.SetIsEnabled(false)

			if responseType == "success" {
				return nil
			}

			if responseType == "error" {
				return fmt.Errorf("error")
			}

			panic("Invalid response type")
		}

		repl := NewREPL(
			WithCommandExecutor(mockedCommandExecutor),
			WithScanner(scanner),
		)

		stdout := utils.NewPrintStorage()
		output = stdout.Capture(func() { repl.RunREPL() })

		userInput = scanner.Text()

		return executedCommand, userInput, output
	}

	t.Run("should execute command and prompt for input until user exits", func(t *testing.T) {
		if executedCommand, expectedCommand, _ := setup("success"); *executedCommand != expectedCommand {
			t.Errorf("Executed command does not match expected command. Got: %s, Expected: %s", *executedCommand, expectedCommand)
		}
	})

	t.Run("should print error message if command execution fails", func(t *testing.T) {
		expectedOutput := "Please try again."
		if _, _, output := setup("error"); !strings.Contains(output, expectedOutput) {
			t.Errorf("Output does not contain expected '%s', but instead got '%s'", expectedOutput, output)
		}
	})
}
