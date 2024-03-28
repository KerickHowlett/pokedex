package repl

import (
	"bufio"
	"fmt"
	"os"
)

// REPL is a struct that provides function dependencies for starting a
// Read-Eval-Print Loop (REPL).
//
// The PrintPrompt field is a function that is called to print the REPL prompt.
//
// The PrintNewLine field is a function that is called to print an empty line in
// the REPL.
//
// The CommandExecutor field is a function that is called to execute commands in
// the REPL.
//
// The Scanner field is an interface that is used to read user input during the
// REPL session.
//
// The ParseInput field is a function that is called to sanitize the user
// input before processing it in the REPL.
type REPL struct {
	CommandExecutor func(string) error
	ParseInput      func(string) []string
	PrintNewLine    func()
	PrintPrompt     func()
	Scanner         scanner
}

// REPLOption is a function type that represents options for configuring a REPL.
type REPLOption func(*REPL)

// scanner is an interface that defines the behavior of a scanner.
//
// It is used to read user input during the REPL session.
//
// The Scan method is used to read the next token from the input.
//
// The Text method is used to return the most recent token read from the input.
type scanner interface {
	Scan() bool
	Text() string
}

// defaultParseInput is the default ParseInput function for the REPL.
func defaultParseInput(input string) []string {
	return []string{input}
}

// defaultPrintNewLine is the default PrintNewLine function for the REPL.
func defaultPrintNewLine() {
	fmt.Println()
}

// defaultPrintPrompt is the default PrintPrompt function for the REPL.
func defaultPrintPrompt() {
	fmt.Print("pokedex > ")
}

// NewREPL creates a new instance of the REPL struct with the provided options.
//
// It accepts variadic REPLOption as parameters.
//
// The returned REPL object is used to start the Read-Eval-Print Loop (REPL).
//
// Parameters:
//   - options: The options to configure the REPL.
//
// Returns:
//   - A new instance of the REPL struct.
//
// Example usage:
//
//	repl := NewREPL(WithCommandExecutor(commandExecutor))
//	repl.StartREPL()
//
// Now the repl can be used to start the REPL.
func NewREPL(options ...REPLOption) *REPL {
	repl := &REPL{
		ParseInput:   defaultParseInput,
		PrintNewLine: defaultPrintNewLine,
		PrintPrompt:  defaultPrintPrompt,
		Scanner:      bufio.NewScanner(os.Stdin),
	}

	for _, option := range options {
		option(repl)
	}

	if repl.CommandExecutor == nil {
		panic("[ERROR] CommandExecutor is required.")
	}

	return repl
}
