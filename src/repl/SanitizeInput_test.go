package repl

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestSanitizeInputByTrimming(t *testing.T) {
	fmt.Println("SanitizeInput() should trim input of any superfluous whitespace.")

	const input = "   hello world   "
	const expected = "hello world"

	if result := strings.Join(SanitizeInput(input), " "); result != expected {
		t.Errorf("SanitizeInput() = %v, want %v", result, expected)
	}
}

func TestSanitizeInputByTransformingToLowercase(t *testing.T) {
	fmt.Println("SanitizeInput() should transform input to lowercase.")

	const input = "HELLO WORLD"
	const expected = "hello world"

	if result := strings.Join(SanitizeInput(input), " "); result != expected {
		t.Errorf("SanitizeInput() = %v, want %v", result, expected)
	}
}

func TestSanitizeInputByExplodingToWordArray(t *testing.T) {
	fmt.Println("SanitizeInput() should copy each word of the argued string into an array.")

	const input = "Hello World"
	expected := []string{"hello", "world"}

	if result := SanitizeInput(input); !reflect.DeepEqual(result, expected) {
		t.Errorf("SanitizeInput() = %v, want %v", result, expected)
	}
}
