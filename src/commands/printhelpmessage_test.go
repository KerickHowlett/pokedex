package commands

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestPrintHelpMessage(t *testing.T) {
	output := setupTestBed(t)
	expected := getExpectedOutputFixture()

	if output != expected {
		t.Errorf("PrintHelpMessage() = %q, want %q", output, expected)
	}
}

// @SECTION: Helper Functions

// @TODO: Mock `GetCommands()` to completely isolate this unit test.
func getExpectedOutputFixture() string {
	expected := "Welcome to the Pokedex!\n"
	expected += "\n" // Empty Line

	expected += "Pokedex Commands\n"

	for _, command := range GetCommands() {
		expected += fmt.Sprintf("%s: %s\n", command.Name, command.Description)
	}

	return expected
}

func getResponseFromPipe(readPipe, writePipe *os.File) string {
	if writePipe != nil {
		writePipe.Close()
	}

	var buf bytes.Buffer
	io.Copy(&buf, readPipe)

	return buf.String()
}

func mockStdout(writePipe *os.File) (savedOriginalStdOut *os.File) {
	savedOriginalStdOut = os.Stdout

	os.Stdout = writePipe // Override stdout with custom write pipe.

	return savedOriginalStdOut
}

func performTest(t *testing.T) {
	if err := PrintHelpMessage(); err != nil {
		t.Errorf("PrintHelpMessage() returned an error: %v", err)
	}
}

func restoreMockedStdout(originalStdout *os.File) {
	os.Stdout = originalStdout
}

func setupTestBed(t *testing.T) string {
	read, write, _ := os.Pipe()

	originalStdout := mockStdout(write)
	defer restoreMockedStdout(originalStdout)

	performTest(t)

	response := getResponseFromPipe(read, write)

	return response
}
