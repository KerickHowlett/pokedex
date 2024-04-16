package repl

import (
	"testing"

	mc "test_tools/mocks/command_executor"
	ms "test_tools/mocks/scanner"
	"test_tools/utils"
)

func TestWithCommandExecutor(t *testing.T) {
	setup := func() func(command string, args []string) error { // Modify the function signature
		repl := NewREPL(WithCommandExecutor(mc.MockedCommandExecutor))
		return repl.execute
	}

	t.Run("should set the Fields field of the Sanitizer struct", func(t *testing.T) {
		execute := setup()
		utils.ExpectSameEntity(t, execute, mc.MockedCommandExecutor, "execute")
	})
}

func TestWithPrompt(t *testing.T) {
	r := &REPL{}
	const prompt = "> "

	t.Run("should set the prompt field in REPL.", func(t *testing.T) {
		WithPrompt(prompt)(r)

		if r.prompt != prompt {
			t.Errorf("Expected r.prompt to be %q, but got %q", prompt, r.prompt)
		}
	})
}

func TestWithScanner(t *testing.T) {
	setup := func() (repl *REPL, scanner *ms.MockScanner) {
		scanner = ms.NewMockScanner()
		repl = NewREPL(
			WithCommandExecutor(mc.MockedCommandExecutor),
			WithScanner(scanner),
		)

		return repl, scanner
	}

	t.Run("should set the Scanner field of the REPL struct", func(t *testing.T) {
		repl, scanner := setup()
		utils.ExpectSameEntity(t, repl.scanner, scanner, "Scanner")
	})
}
