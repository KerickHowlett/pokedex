package repl

import "testing"

func TestWithPrompt(t *testing.T) {
	r := &REPL{}
	const prompt = "Test Prompt"

	t.Run("should set the prompt field in REPL.", func(t *testing.T) {
		WithPrompt(prompt)(r)

		if r.prompt != prompt {
			t.Errorf("Expected r.prompt to be %q, but got %q", prompt, r.prompt)
		}
	})
}
