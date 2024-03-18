package toolchain

import (
	"fmt"
	"testing"

	"github.com/KerickHowlett/pokedexcli/tests/mocks"
	"github.com/KerickHowlett/pokedexcli/tests/utils"
)

func TestRunCommand_PrintHelpMessage(t *testing.T) {
	printout := utils.PrintStorage{}
	toolchain := NewToolchain()

	printout.Capture(toolchain.PrintHelpMessage)

	expectedOutput := createExpectedPrintHelpMessageOutput(toolchain)
	actualOutput := printout.Capture(toolchain.PrintHelpMessage)

	if actualOutput != expectedOutput {
		t.Errorf("Expected output to be the following:\n %q\n\n Got the following instead:\n %q", expectedOutput, actualOutput)
	}
}

func TestRunCommand_ExecuteSelectedCommand(t *testing.T) {
	mockCommand := &mocks.MockCommand{Name: "Command 1", Description: "Command 1 help message."}
	toolchain := NewToolchain(WithCommand(mockCommand))

	err := toolchain.RunCommand("Command 1")

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestRunCommand_InvalidCommand(t *testing.T) {
	toolchain := NewToolchain()

	err := toolchain.RunCommand("InvalidCommand")

	expectedError := fmt.Errorf("[ERROR] Command 'InvalidCommand' is not valid")
	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error to be: %v, but got: %v", expectedError, err)
	}
}

func TestRunCommand_EmptySelection(t *testing.T) {
	toolchain := NewToolchain()

	err := toolchain.RunCommand("")

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}
