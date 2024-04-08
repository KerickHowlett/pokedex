package command

import (
	"fmt"
	"testing"

	m "test_tools/mocks/command"
	"test_tools/utils"
)

func TestPrintHelp(t *testing.T) {
	setup := func() (command *m.MockCommand, printout *utils.PrintStorage) {
		command = m.NewMockCommand()
		printout = utils.NewPrintStorage()

		return command, printout
	}

	t.Run("should print the command name and description in the proper format.", func(t *testing.T) {
		command, printout := setup()
		output := printout.Capture(command.PrintHelp)

		expectedOutput := fmt.Sprintf("%s: %s\n", command.GetName(), command.GetDescription())
		if output != expectedOutput {
			t.Errorf("Expected output to be %q, but got %q", expectedOutput, output)
		}
	})
}
