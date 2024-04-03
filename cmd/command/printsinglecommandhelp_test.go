package command

import (
	"fmt"
	"testing"

	m "testtools/mocks/command"
	"testtools/utils"
)

func TestPrintSingleCommandHelp(t *testing.T) {
	printout := utils.NewPrintStorage()

	t.Run("should print the command name and description in the proper format.", func(t *testing.T) {
		command := m.NewMockCommand()
		output := printout.Capture(command.PrintHelp)

		expectedOutput := fmt.Sprintf("%s: %s\n", command.GetName(), command.GetDescription())
		if output != expectedOutput {
			t.Errorf("Expected output to be %q, but got %q", expectedOutput, output)
		}
	})
}
