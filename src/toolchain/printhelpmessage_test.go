package toolchain

import (
	"fmt"
	"testing"

	utils "github.com/KerickHowlett/pokedexcli/testing_utils"
	c "github.com/KerickHowlett/pokedexcli/toolchain/command"
)

// @SECTION: Unit Test Cases

func TestPrintHelpMessage(t *testing.T) {
	printout := utils.PrintStorage{}

	command1 := &c.MockCommand{Name: "Command 1", Description: "Command 1 help message."}
	command2 := &c.MockCommand{Name: "Command 2", Description: "Command 2 help message."}
	toolchain := NewToolchain(WithCommand(command1), WithCommand(command2))

	expectedOutput := createExpectedPrintHelpMessageOutput(toolchain)

	output := printout.Capture(toolchain.PrintHelpMessage)

	if output != expectedOutput {
		t.Errorf("Expected output to be the following:\n %q\n\n Got the following instead:\n %q", expectedOutput, output)
	}
}

// @SECTION: Helper Functions

func createExpectedPrintHelpMessageOutput(toolchain *Toolchain) string {
	output := "Welcome to the Pokedex!\n"
	output += "\n" // Empty Line
	output += "Pokedex Commands\n"
	output += "--------------------\n"
	output += "\n" // Empty Line

	for _, command := range *toolchain.commands {
		output += fmt.Sprintf("%s: %s\n", (*command).GetName(), (*command).GetDescription())
	}

	return output
}
