package repl

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestPrintPrompt(t *testing.T) {
	fmt.Println("PrintPrompt() should print the REPL prompt to stdout.")

	const expected = "pokedex > "
	if actual := setupPrintPromptTestBed(t); actual != expected {
		t.Errorf("Expected prompt '%s', but got '%s'", expected, actual)
	}
}

func setupPrintPromptTestBed(t *testing.T) string {
	// Capture stdout from IO Pipe.
	read, write, _ := os.Pipe()
	originalStdout := os.Stdout
	os.Stdout = write
	defer func() { os.Stdout = originalStdout }()

	PrintPrompt() // Perform the test.

	// Copy captured value to set as the `response`.
	write.Close()
	var buf bytes.Buffer
	io.Copy(&buf, read)
	response := buf.String()

	return response
}
