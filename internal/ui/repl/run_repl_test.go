package repl

import (
	"fmt"
	"reflect"
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

func TestParseInput(t *testing.T) {
	runParseInputTest := func(input string, expected []string) {
		repl := &REPL{}
		response := repl.parseInput(input)
		if !reflect.DeepEqual(response, expected) {
			t.Errorf("Expected %v, but instead got %v\n", expected, response)
		}
	}

	t.Run("should trim leading and trailing spaces", func(t *testing.T) {
		runParseInputTest("  Pikachu  ", []string{"pikachu"})
	})

	t.Run("should convert the input to lowercase", func(t *testing.T) {
		runParseInputTest("PIKACHU", []string{"pikachu"})
	})

	t.Run("should split each word of input into a slice of words.", func(t *testing.T) {
		runParseInputTest("Team Rocket", []string{"team", "rocket"})
	})

	t.Run("should return an empty string within a slice if no input is provided", func(t *testing.T) {
		runParseInputTest("", []string{""})
	})
}
func TestREPL_printNewLine(t *testing.T) {
	setup := func() (output string) {
		stdout := utils.NewPrintStorage()
		output = stdout.Capture(REPL{}.printNewLine)

		return output
	}

	t.Run("should print out an empty new line", func(t *testing.T) {
		if output := setup(); output != "\n" {
			t.Errorf("Expected output to be '\\n'. Got '%s'", output)
		}
	})
}

func TestREPL_printPrompt(t *testing.T) {
	r := REPL{prompt: " >"}
	stdout := utils.NewPrintStorage()

	t.Run("should print the set prompt to stdout", func(t *testing.T) {
		output := stdout.Capture(r.printPrompt)

		if output != r.prompt {
			t.Errorf("Expected output to be %q. Got %q", r.prompt, output)
		}
	})
}
