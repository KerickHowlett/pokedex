package repl

type REPL struct {
	execute func(string) error
	Scanner scanner
}
