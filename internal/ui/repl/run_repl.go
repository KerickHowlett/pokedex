package repl

import (
	"fmt"
	"strings"
)

// RunREPL starts the Read-Eval-Print Loop (REPL) for the REPL struct.
// It continuously prompts the user for input, sanitizes the input, executes the
// command, and prints the result. The loop continues until the user exits the
// REPL.
//
// Example usage:
//
//	repl := NewREPL(WithCommandExecutor(commandExecutor))
//	repl.RunREPL()
func (r *REPL) RunREPL() {
	r.printPrompt()

	for r.scanner.Scan() {
		r.printNewLine()

		rawInput := r.scanner.Text()
		parsedInput := r.parseInput(rawInput)

		command, args := parsedInput[0], parsedInput[1:]
		if err := r.execute(command, args); err != nil {
			fmt.Println(err)
			fmt.Println("Please try again.")
		}

		r.printNewLine()
		r.printPrompt()
	}
}

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

	if len(words) == 0 {
		noEnteredCommands := []string{""}
		return noEnteredCommands
	}

	return words
}

// printNewLine prints a new line to the console.
//
// Example usage:
//
//	repl := NewREPL()
//	repl.printNewLine()
func (r REPL) printNewLine() {
	fmt.Println()
}

// printPrompt prints the REPL prompt to the console.
//
// Example usage:
//
//	repl := NewREPL()
//	repl.printPrompt()
func (r REPL) printPrompt() {
	fmt.Print(r.prompt)
}
