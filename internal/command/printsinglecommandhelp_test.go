package command

import (
	"fmt"
	"testing"

	utils "github.com/KerickHowlett/pokedexcli/internal/testing_utils"
)

func TestPrintSingleCommandHelp(t *testing.T) {
	printout := utils.PrintStorage{}

	mock := &MockCommand{Name: "mock", Description: "This is a mocked command."}
	expectedOutput := fmt.Sprintf("%s: %s\n", mock.GetName(), mock.GetDescription())

	output := printout.Capture(mock.PrintHelp)

	// Compare the actual and expected output
	if output != expectedOutput {
		t.Errorf("Expected output to be %q, but got %q", expectedOutput, output)
	}
}
