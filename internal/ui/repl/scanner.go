package repl

// scanner is an interface that represents a scanner used in the REPL (Read-Eval-Print Loop).
// It provides methods for scanning input and retrieving the scanned text.
//
// Fields:
//   - Scan: Scans the input and returns true if there is more input to be scanned, false otherwise.
//   - Text: Returns the text that has been scanned.
type scanner interface {
	Scan() bool
	Text() string
}
