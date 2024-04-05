package repl

import (
	"testing"
	f "testtools/fixtures"
	"testtools/utils"
)

func TestREPL_printPrompt(t *testing.T) {
	r := REPL{prompt: f.Prompt}
	stdout := utils.NewPrintStorage()

	t.Run("should print the set prompt to stdout", func(t *testing.T) {
		output := stdout.Capture(r.printPrompt)

		if output != r.prompt {
			t.Errorf("Expected output to be %q. Got %q", r.prompt, output)
		}
	})
}
