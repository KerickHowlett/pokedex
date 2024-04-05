package repl

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
