package repl

import "strings"

// This function takes a string and returns a slice of strings after performing
// the necessary sanitization.
func SanitizeInput(input string) []string {
	trimmed := strings.TrimSpace(input)
	lowered := strings.ToLower(trimmed)
	words := strings.Fields(lowered)

	return words
}
