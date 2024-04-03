package repl

import (
	"testing"
	"testtools/utils"
)

func TestREPL_printNewLine(t *testing.T) {
	const emptyNewLine = "\n"

	r := REPL{}
	stdout := utils.NewPrintStorage()

	t.Run("should print out an empty new line", func(t *testing.T) {
		output := stdout.Capture(r.printNewLine)

		if output != emptyNewLine {
			t.Errorf("Expected output to be '\\n'. Got '%s'", output)
		}
	})
}
