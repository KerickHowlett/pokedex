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
// The SanitizeInput field is a function that is called to sanitize the user
// input before processing it in the REPL.
type REPL struct {
	PrintPrompt     func()
	PrintNewLine    func()
	CommandExecutor func(...string) error
	Scanner         Scanner
	SanitizeInput   func(string) []string
}

// REPLOption is a function type that represents options for configuring a REPL.
type REPLOption func(*REPL)

// Scanner is an interface that defines the behavior of a scanner.
//
// It is used to read user input during the REPL session.
//
// The Scan method is used to read the next token from the input.
//
// The Text method is used to return the most recent token read from the input.
type Scanner interface {
	Scan() bool
	Text() string
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
		PrintPrompt:   func() { fmt.Print("pokedex > ") },
		PrintNewLine:  func() { fmt.Println() },
		Scanner:       bufio.NewScanner(os.Stdin),
		SanitizeInput: NewSanitizer().SanitizeInput,
	}

	for _, option := range options {
		option(repl)
	}

	if repl.CommandExecutor == nil {
		panic("[ERROR] CommandExecutor is required.")
	}

	return repl
}

// WithCommandExecutor is a function that sets the command executor for the REPL.
//
// The command executor is responsible for executing commands in the REPL.
//
// It takes a variadic argument of strings representing the command and returns
// an error if any.
//
// This function is used as an option in the REPL constructor to customize the
// behavior of the REPL.
//
// Parameters:
//   - commandExecutor: The function to be called for executing commands in the
//     REPL.
//
// Returns:
//   - An option function that sets the commandExecutor function for the REPL.
//
// Example usage:
//
//	commandExecutor := func(args ...string) error {
//		fmt.Println("Executing command:", args)
//		return nil
//	}
//	repl := NewREPL(WithCommandExecutor(commandExecutor))
//
// Now the commandExecutor function will be used to execute commands in the REPL.
func WithCommandExecutor(commandExecutor func(...string) error) REPLOption {
	return func(r *REPL) {
		r.CommandExecutor = commandExecutor
	}
}

// WithPrintNewLine is a REPLOption function that sets the PrintNewLine function
// of the REPL.
//
// The PrintNewLine function is called to print an empty line in the REPL.
//
// It takes a printEmptyLine function as a parameter and assigns it to the
// PrintNewLine field of the REPL.
//
// Parameters:
//   - printEmptyLine: The function to be called for printing an empty line.
//
// Returns:
//   - An option function that sets the printEmptyLine function for the REPL.
//
// Example usage:
//
//	printEmptyLine := func() {
//		fmt.Println()
//	}
//	repl := NewREPL(WithPrintNewLine(printEmptyLine))
//
// Now the printEmptyLine function will be used to print an empty line in the
// REPL.
func WithPrintNewLine(printEmptyLine func()) REPLOption {
	return func(r *REPL) {
		r.PrintNewLine = printEmptyLine
	}
}

// WithPrintPrompt sets the function to be called for printing the REPL prompt.
// It is an option function that can be used with the REPL struct.
//
// Parameters:
//   - printPrompt: The function to be called for printing the REPL prompt.
//
// Returns:
//   - An option function that sets the printPrompt function for the REPL.
//
// Example usage:
//
//	printPrompt := func() {
//		fmt.Print("pokedex > ")
//	}
//	repl := NewREPL(WithPrintPrompt(printPrompt))
//
// Now the printPrompt function will be used to print the REPL prompt.
func WithPrintPrompt(printPrompt func()) REPLOption {
	return func(r *REPL) {
		r.PrintPrompt = printPrompt
	}
}

// WithSanitizeInput is a REPLOption function that sets the sanitizeInput
// function for the REPL.
//
// The sanitizeInput function takes a string as input and returns a slice of
// strings after sanitizing the input.
//
// It is used to sanitize the user input before processing it in the REPL.
//
// Parameters:
//   - sanitizeInput: The sanitizeInput function to be set for the REPL.
//
// Returns:
//   - An option function that sets the sanitizeInput function for the REPL.
//
// Example usage:
//
//	sanitizeInput := func(input string) []string {
//		return strings.Fields(input)
//	}
//	repl := NewREPL(WithSanitizeInput(sanitizeInput))
//
// Now the sanitizeInput function will be used to sanitize user input during the
// REPL session.
func WithSanitizeInput(sanitizeInput func(string) []string) REPLOption {
	return func(r *REPL) {
		r.SanitizeInput = sanitizeInput
	}
}

// WithScanner sets the scanner for the REPL.
// It is an option function that can be used when creating a new REPL instance.
// The scanner is responsible for reading user input during the REPL session.
//
// Parameters:
//   - scanner: The scanner to be set for the REPL.
//
// Returns:
//   - An option function that sets the scanner for the REPL.
//
// Example usage:
//
//	scanner := bufio.NewScanner(os.Stdin)
//	repl := NewREPL(WithScanner(scanner))
//
// Now the scanner will be used to read user input during the REPL session.
func WithScanner(scanner Scanner) REPLOption {
	return func(r *REPL) {
		r.Scanner = scanner
	}
}

// StartREPL starts the Read-Eval-Print Loop (REPL) for the REPL struct.
// It continuously prompts the user for input, sanitizes the input, executes the
// command, and prints the result. The loop continues until the user exits the
// REPL.
//
// Example usage:
//
//	repl := NewREPL(WithCommandExecutor(commandExecutor))
//	repl.StartREPL()
//
// The StartREPL method is used to start the REPL.
func (r *REPL) StartREPL() {
	r.PrintPrompt()

	for r.Scanner.Scan() {
		userInput := r.Scanner.Text()
		command := r.SanitizeInput(userInput)[0]

		r.PrintNewLine()
		r.CommandExecutor(command)

		r.PrintNewLine()
		r.PrintPrompt()
	}
}
