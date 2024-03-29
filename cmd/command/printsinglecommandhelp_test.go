package command

import (
	"fmt"
	"testing"

	m "internal/tests/mocks"
	"internal/tests/utils"
)

func TestPrintSingleCommandHelp(t *testing.T) {
	printout := utils.NewPrintStorage()

	command := &m.MockCommand{Name: "mock", Description: "This is a mocked command."}
	expectedOutput := fmt.Sprintf("%s: %s\n", command.GetName(), command.GetDescription())

	output := printout.Capture(command.PrintHelp)

	// Compare the actual and expected output
	if output != expectedOutput {
		t.Errorf("Expected output to be %q, but got %q", expectedOutput, output)
	}
}
