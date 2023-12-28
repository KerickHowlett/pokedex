package repl

import "strings"

func SanitizeInput(input string) []string {
	trimmed := strings.TrimSpace(input)
	lowered := strings.ToLower(trimmed)
	words := strings.Fields(lowered)

	return words
}
