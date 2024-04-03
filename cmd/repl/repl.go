package repl

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
