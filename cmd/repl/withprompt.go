package repl

func WithPrompt(prompt string) REPLOption {
	return func(r *REPL) {
		r.prompt = prompt
	}
}
