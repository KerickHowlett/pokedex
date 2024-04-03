package repl

import (
	"testing"
	f "testtools/fixtures"
)

func TestWithPrompt(t *testing.T) {
	r := &REPL{}

	t.Run("should set the prompt field in REPL.", func(t *testing.T) {
		WithPrompt(f.Prompt)(r)

		if r.prompt != f.Prompt {
			t.Errorf("Expected r.prompt to be %q, but got %q", f.Prompt, r.prompt)
		}
	})
}
