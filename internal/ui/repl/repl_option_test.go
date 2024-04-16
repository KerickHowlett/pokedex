package repl

import (
	"testing"

	mc "test_tools/mocks/command_executor"
	ms "test_tools/mocks/scanner"
	"test_tools/utils"
)

func TestWithErrorMessagePrompt(t *testing.T) {
	runWithErrorMessagePromptTest := func() (repl *REPL, errorMessagePrompt string) {
		repl = &REPL{}
		errorMessagePrompt = "Error Prompt"
		WithErrorMessagePrompt(errorMessagePrompt)(repl)
		return repl, errorMessagePrompt
	}
	t.Run("should set the errorMessagePrompt field in REPL.", func(t *testing.T) {
		if repl, errorMessagePrompt := runWithErrorMessagePromptTest(); repl.errorMessagePrompt != errorMessagePrompt {
			t.Errorf("Expected r.errorMessagePrompt to be %q, but got %q", errorMessagePrompt, repl.errorMessagePrompt)
		}
	})
}

func TestWithCommandExecutor(t *testing.T) {
	setup := func() CommandExecutor {
		repl := &REPL{}
		WithCommandExecutor(mc.MockedCommandExecutor)(repl)
		return repl.execute
	}
	t.Run("should set the Fields field of the Sanitizer struct", func(t *testing.T) {
		execute := setup()
		utils.ExpectSameEntity(t, execute, mc.MockedCommandExecutor, "execute")
	})
}

func TestWithPrompt(t *testing.T) {
	runWithPromptTest := func() (repl *REPL, prompt string) {
		repl = &REPL{}
		prompt = "*"
		WithPrompt(prompt)(repl)
		return repl, prompt
	}
	t.Run("should set the prompt field in REPL.", func(t *testing.T) {
		if repl, prompt := runWithPromptTest(); repl.prompt != prompt {
			t.Errorf("Expected REPL.prompt to be %q, but got %q", prompt, repl.prompt)
		}
	})
}

func TestWithScanner(t *testing.T) {
	setup := func() (repl *REPL, scanner *ms.MockScanner) {
		repl = &REPL{}
		scanner = ms.NewMockScanner()
		WithScanner(scanner)(repl)
		return repl, scanner
	}

	t.Run("should set the Scanner field of the REPL struct", func(t *testing.T) {
		repl, scanner := setup()
		utils.ExpectSameEntity(t, repl.scanner, scanner, "Scanner")
	})
}
