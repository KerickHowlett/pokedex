package repl

import "strings"

func SanitizeInput(input string) string {
	trimmedInput := strings.TrimSpace(input)

	return strings.ToLower(trimmedInput)
}
