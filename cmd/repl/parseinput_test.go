package repl

import (
	"fmt"
	"testing"
)

func TestParseInput(t *testing.T) {
	fmt.Println("parseInput should return a sanitized string")

	repl := &REPL{}

	response := repl.parseInput("  Hello World  ")

	const expected string = "hello"
	if response == nil || response[0] != expected {
		t.Errorf("Expected %s, but instead got %s\n", expected, response)
	}
}
