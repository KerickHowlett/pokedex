package parser

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestDefaultNewSanitizer(t *testing.T) {
	fmt.Println("NewSanitizer should return a Sanitizer with default values")
	sanitizer := NewParser()

	expectSameEntity(t, sanitizer.Fields, strings.Fields, "Fields")
	expectSameEntity(t, sanitizer.ToLower, strings.ToLower, "ToLower")
	expectSameEntity(t, sanitizer.TrimSpace, strings.TrimSpace, "TrimSpace")
}

func TestWithFields(t *testing.T) {
	fmt.Println("WithFields should set the Fields field of the Sanitizer struct")
	sanitizer := NewParser(WithFields(stringSliceReturnMock))

	expectSameEntity(t, sanitizer.Fields, stringSliceReturnMock, "Fields")
}

func TestWithToLower(t *testing.T) {
	fmt.Println("WithToLower should set the ToLower field of the Sanitizer struct")
	sanitizer := NewParser(WithToLower(stringReturnMock))

	expectSameEntity(t, sanitizer.ToLower, stringReturnMock, "ToLower")
}

func TestWithTrimSpace(t *testing.T) {
	fmt.Println("WithTrimSpace should set the TrimSpace field of the Sanitizer struct")
	sanitizer := NewParser(WithTrimSpace(stringReturnMock))

	expectSameEntity(t, sanitizer.TrimSpace, stringReturnMock, "TrimSpace")
}

func TestParseInput(t *testing.T) {
	fmt.Println("ParseInput should return a sanitized string")

	const expected string = "hello"
	sanitizer := NewParser()

	response := sanitizer.ParseInput("  Hello World  ")

	if response == nil || response[0] != expected {
		t.Errorf("Expected %s, but instead got %s\n", expected, response)
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

func expectSameEntity(t *testing.T, actual any, expected any, fieldName string) {
	if reflect.ValueOf(actual).Pointer() == reflect.ValueOf(expected).Pointer() {
		t.Errorf("Expected %s to be set with the argued function, but instead got %v\n", fieldName, actual)
	}
}
