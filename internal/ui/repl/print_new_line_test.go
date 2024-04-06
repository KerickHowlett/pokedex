package repl

import (
	"testing"

	"test_tools/utils"
)

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
