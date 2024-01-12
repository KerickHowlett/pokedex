package repl

import "strings"

// This function takes a string and returns a slice of strings after performing
// the necessary sanitization.
func SanitizeInput(input string) []string {
	trimmedInput := strings.TrimSpace(input)
	lowercasedInput := strings.ToLower(trimmedInput)

	words := strings.Fields(lowercasedInput)

	return words
}
