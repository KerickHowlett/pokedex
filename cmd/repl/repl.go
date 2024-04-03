package repl

type REPL struct {
	execute func(string) error
	scanner scanner
}
