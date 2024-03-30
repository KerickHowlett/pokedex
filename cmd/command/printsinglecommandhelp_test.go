package command

import (
	"fmt"
	"testing"

	m "internal/tests/mocks/command"
	"internal/tests/utils"
)

func TestPrintSingleCommandHelp(t *testing.T) {
	printout := utils.NewPrintStorage()

	command := m.NewMockCommand()
	output := printout.Capture(command.PrintHelp)

	expectedOutput := fmt.Sprintf("%s: %s\n", command.GetName(), command.GetDescription())
	if output != expectedOutput {
		t.Errorf("Expected output to be %q, but got %q", expectedOutput, output)
	}
}
