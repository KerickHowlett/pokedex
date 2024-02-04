package parser

import (
	"fmt"
	"strings"
	"testing"
)

func TestDefaultNewSanitizer(t *testing.T) {
	fmt.Println("NewSanitizer should return a Sanitizer with default values")
	sanitizer := NewParser()

	expectedFunction(t, sanitizer.Fields, strings.Fields, "Expected Fields to be set with strings.Fields, but found nil\n.")
	expectedFunction(t, sanitizer.ToLower, strings.ToLower, "Expected ToLower to be set with strings.Fields, but found nil\n.")
	expectedFunction(t, sanitizer.TrimSpace, strings.TrimSpace, "Expected TrimSpace to be set with strings.Fields, but found nil\n.")
}

func TestWithFields(t *testing.T) {
	fmt.Println("WithFields should set the Fields field of the Sanitizer struct")
	sanitizer := NewParser(WithFields(stringSliceReturnMock))

	if fmt.Sprintf("%p", sanitizer.Fields) != fmt.Sprintf("%p", stringSliceReturnMock) {
		t.Errorf("Expected Fields to be set, got nil")
	}
}

type FuncMock func(argument string) string

func TestWithToLower(t *testing.T) {
	fmt.Println("WithToLower should set the ToLower field of the Sanitizer struct")
	sanitizer := NewParser(WithToLower(stringReturnMock))

	if fmt.Sprintf("%p", sanitizer.ToLower) != fmt.Sprintf("%p", stringReturnMock) {
		t.Errorf("Expected toLowerMock to be set, got nil")
	}
}

func TestWithTrimSpace(t *testing.T) {
	fmt.Println("WithTrimSpace should set the TrimSpace field of the Sanitizer struct")
	actual := NewParser(WithTrimSpace(stringReturnMock)).TrimSpace

	if fmt.Sprintf("%p", actual) != fmt.Sprintf("%p", stringReturnMock) {
		t.Errorf("Expected trimSpace to be set, got nil")
	}
}

func TestParseInput(t *testing.T) {
	fmt.Println("ParseInput should return a sanitized string")

	const expected string = "hello"
	sanitizer := NewParser()

	response := sanitizer.ParseInput("  Hello World  ")

	if response == nil || response[0] != expected {
		t.Errorf("Expected %s, got %s\n", expected, response)
	}
}

// @SECTION: MOCKS

func stringSliceReturnMock(argument string) []string {
	return []string{argument}
}

func stringReturnMock(argument string) string {
	return argument
}

// @SECTION: ASSERTERS

func expectedFunction(t *testing.T, actual interface{}, expected interface{}, message string) {
	if fmt.Sprintf("%p", actual) != fmt.Sprintf("%p", expected) {
		t.Errorf(message)
	}
}
