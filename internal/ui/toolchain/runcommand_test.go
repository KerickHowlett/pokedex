package toolchain

import (
	"testing"

	m "testtools/mocks/command"
	"testtools/utils"
)

func TestRunCommand(t *testing.T) {
	toolchain := NewToolchain(WithCommand(m.NewMockCommand()))

	helpCommandSetup := func(toolchain *Toolchain) (printout *utils.PrintStorage, expectedOutput string) {
		printout = utils.NewPrintStorage()
		expectedOutput = getExpectedPrintHelpMessage(toolchain)

		return printout, expectedOutput
	}

	t.Run("should print the help message if the command is not found.", func(t *testing.T) {
		printout, expectedOutput := helpCommandSetup(toolchain)
		if output := printout.Capture(func() { toolchain.RunCommand("help") }); output != expectedOutput {
			t.Errorf("Expected output to be the following:\n %q\n\n Got the following instead:\n %q", expectedOutput, output)
		}
	})

	t.Run("should execute selected command without returning an error.", func(t *testing.T) {
		if err := toolchain.RunCommand("mock"); err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
	})

	t.Run("should return an error if entered command is invalid.", func(t *testing.T) {
		const expectedError = "[ERROR] Command 'InvalidCommand' is not valid"
		if err := toolchain.RunCommand("InvalidCommand"); err.Error() != expectedError {
			t.Errorf("Expected error to be: %v, but got: %v", expectedError, err)
		}
	})

	t.Run("should not do anything when no command is entered.", func(t *testing.T) {
		if err := toolchain.RunCommand(""); err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
	})
}
