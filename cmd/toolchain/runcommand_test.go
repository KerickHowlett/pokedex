package toolchain

import (
	"fmt"
	"testing"

	m "internal/tests/mocks/command"
	"internal/tests/utils"
)

func TestRunCommand_PrintHelpMessage(t *testing.T) {
	toolchain := createMockToolchain()

	printout := utils.NewPrintStorage()
	printout.Capture(toolchain.PrintHelpMessage)

	expectedOutput := createExpectedPrintHelpMessageOutput(toolchain)
	actualOutput := printout.Capture(toolchain.PrintHelpMessage)

	if actualOutput != expectedOutput {
		t.Errorf("Expected output to be the following:\n %q\n\n Got the following instead:\n %q", expectedOutput, actualOutput)
	}
}

func TestRunCommand_ExecuteSelectedCommand(t *testing.T) {
	toolchain := createMockToolchain()

	err := toolchain.RunCommand("mock")

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestRunCommand_InvalidCommand(t *testing.T) {
	toolchain := createMockToolchain()

	err := toolchain.RunCommand("InvalidCommand")

	expectedError := fmt.Errorf("[ERROR] Command 'InvalidCommand' is not valid")
	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error to be: %v, but got: %v", expectedError, err)
	}
}

func TestRunCommand_EmptySelection(t *testing.T) {
	toolchain := createMockToolchain()

	err := toolchain.RunCommand("")

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func createMockToolchain() *Toolchain {
	command := m.NewMockCommand()
	return NewToolchain(WithCommand(command))
}
