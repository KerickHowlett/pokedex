package commands

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestPrintHelp(t *testing.T) {
	command := Command{
		Name:        "test",
		Description: "This is a test command",
		Execute:     func() error { return nil },
	}
	const expected string = "test: This is a test command\n"

	if response := performTest(command, t); response != expected {
		t.Errorf("PrintHelp() = %v, want %v", response, expected)
	}
}

func performTest(command Command, t *testing.T) string {
	// Capture stdout from IO Pipe.
	read, write, _ := os.Pipe()
	originalStdout := os.Stdout
	os.Stdout = write
	defer func() { os.Stdout = originalStdout }()

	// Perform the test.
	command.PrintHelp()

	// Copy captured value to set as the `response`.
	write.Close()
	var buf bytes.Buffer
	io.Copy(&buf, read)
	response := buf.String()

	return response
}
