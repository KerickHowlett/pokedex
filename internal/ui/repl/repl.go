package repl

type REPL struct {
	execute func(string) error
	prompt  string
	scanner scanner
}
