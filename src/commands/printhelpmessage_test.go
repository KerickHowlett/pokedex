package commands

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

// @SECTION: Unit Test Cases

func TestPrintHelpMessage(t *testing.T) {
	fmt.Println("Given a list of commands, PrintHelpMessage() should print out the help message.")

	output := setupTestBed(PrintMessageCommandsFixture, t)
	expected := getExpectedOutputFixture(PrintMessageCommandsFixture)

	if output != expected {
		t.Errorf("PrintHelpMessage() = %q, want %q", output, expected)
	}
}

// @SECTION: Helper Fixtures & Functions

var PrintMessageCommandsFixture Commands = Commands{
	"test": {
		Name:        "test",
		Description: "This is a fake command.",
		Execute:     func() error { return nil },
	},
}

func getExpectedOutputFixture(commands Commands) string {
	expected := "Welcome to the Pokedex!\n"
	expected += "\n" // Empty Line
	expected += "Pokedex Commands\n"

	for _, command := range commands {
		expected += fmt.Sprintf("%s: %s\n", command.Name, command.Description)
	}

	return expected
}

func setupTestBed(commands Commands, t *testing.T) string {
	// Capture stdout from IO Pipe.
	read, write, _ := os.Pipe()
	originalStdout := os.Stdout
	os.Stdout = write
	defer func() { os.Stdout = originalStdout }()

	// Perform the test.
	if err := PrintHelpMessage(commands); err != nil {
		t.Errorf("PrintHelpMessage() returned an error: %v", err)
	}

	// Copy captured value to set as the `response`.
	write.Close()
	var buf bytes.Buffer
	io.Copy(&buf, read)
	response := buf.String()

	return response
}
