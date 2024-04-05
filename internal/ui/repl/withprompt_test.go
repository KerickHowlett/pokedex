package repl

import "testing"

func TestWithPrompt(t *testing.T) {
	r := &REPL{}

	t.Run("should set the prompt field in REPL.", func(t *testing.T) {
		WithPrompt(defaultPrompt)(r)

		if r.prompt != defaultPrompt {
			t.Errorf("Expected r.prompt to be %q, but got %q", defaultPrompt, r.prompt)
		}
	})
}
