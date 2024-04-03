package repl

import "strings"

// parseInput sanitizes the input string by trimming leading and trailing
// spaces, converting it to lowercase, and splitting it into individual words.
//
// It returns a slice of strings representing the sanitized words.
//
// Parameters:
//   - input: The input string to be sanitized.
//
// Returns:
//   - A slice of strings representing the sanitized words.
//
// Example usage:
//
//	sanitizedOutput := sanitizer.parseInput("  HELLO  ")
func (r *REPL) parseInput(input string) (words []string) {
	trimmedInput := strings.TrimSpace(input)
	lowercasedInput := strings.ToLower(trimmedInput)
	words = strings.Fields(lowercasedInput)

	return words
}
