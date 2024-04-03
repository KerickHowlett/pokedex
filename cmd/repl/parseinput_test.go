package repl

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	repl := &REPL{}
	const emptyString = ""

	runParseInputTest := func(input string, expected []string) {
		response := repl.parseInput(input)
		if !reflect.DeepEqual(response, expected) {
			t.Errorf("Expected %v, but instead got %v\n", expected, response)
		}
	}

	t.Run("should trim leading and trailing spaces", func(t *testing.T) {
		runParseInputTest("  Pikachu  ", []string{"pikachu"})
	})

	t.Run("should convert the input to lowercase", func(t *testing.T) {
		runParseInputTest("PIKACHU", []string{"pikachu"})
	})

	t.Run("should split each word of input into a slice of words.", func(t *testing.T) {
		runParseInputTest("Team Rocket", []string{"team", "rocket"})
	})

	t.Run("should return an empty string within a slice if no input is provided", func(t *testing.T) {
		runParseInputTest(emptyString, []string{emptyString})
	})
}
