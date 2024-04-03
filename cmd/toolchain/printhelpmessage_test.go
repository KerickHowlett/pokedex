package toolchain

import (
	"fmt"
	"testing"

	m "testtools/mocks/command"
	"testtools/utils"
)

func TestPrintHelpMessage(t *testing.T) {
	stdout := utils.NewPrintStorage()

	setup := func() (toolchain *Toolchain, expectedOutput string) {
		command1 := m.NewMockCommand(
			m.WithName("Command 1"),
			m.WithDescription("Command 1 help message."),
		)
		command2 := m.NewMockCommand(
			m.WithName("Command 2"),
			m.WithDescription("Command 2 help message."),
		)
		toolchain = NewToolchain(
			WithCommand(command1),
			WithCommand(command2),
		)

		expectedOutput = createExpectedPrintHelpMessageOutput(toolchain)

		return toolchain, expectedOutput
	}

	t.Run("should print the help message for the toolchain.", func(t *testing.T) {
		toolchain, expectedOutput := setup()

		output := stdout.Capture(toolchain.PrintHelpMessage)

		if output != expectedOutput {
			t.Errorf("Expected output to be the following:\n %q\n\n Got the following instead:\n %q", expectedOutput, output)
		}
	})
}

func createExpectedPrintHelpMessageOutput(toolchain *Toolchain) string {
	const emptyLine = "\n"

	expectedOutput := "Welcome to the Pokedex!\n"
	expectedOutput += emptyLine
	expectedOutput += "Pokedex Commands\n"
	expectedOutput += "--------------------\n"
	expectedOutput += emptyLine

	for _, command := range *toolchain.commands {
		name, description := (*command).GetName(), (*command).GetDescription()
		expectedOutput += fmt.Sprintf("%s: %s\n", name, description)
	}

	return expectedOutput
}
