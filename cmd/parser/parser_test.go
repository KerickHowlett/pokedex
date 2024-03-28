package parser

import (
	"fmt"
	"testing"
)

func TestParseInput(t *testing.T) {
	fmt.Println("ParseInput should return a sanitized string")

	const expected string = "hello"

	response := ParseInput("  Hello World  ")

	if response == nil || response[0] != expected {
		t.Errorf("Expected %s, but instead got %s\n", expected, response)
	}
}
