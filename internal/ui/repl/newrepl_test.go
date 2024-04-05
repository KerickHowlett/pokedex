package repl

import (
	"bufio"
	"testing"

	"testtools/utils"
)

// @SECTION: TEST CASES

func TestNewREPL(t *testing.T) {
	t.Run("REPL created successfully", func(t *testing.T) {
		const defaultPrompt = "> "
		mockedCommandExecutor := func(args string) error { return nil }

		repl := NewREPL(WithCommandExecutor(mockedCommandExecutor))

		t.Run("with required argument(s)", func(t *testing.T) {
			t.Run("should have the provided command executor", func(t *testing.T) {
				utils.ExpectSameEntity(t, repl.execute, mockedCommandExecutor, "execute")
			})
		})

		t.Run("with the default optional argument(s)", func(t *testing.T) {
			t.Run("should have the preset default scanner", func(t *testing.T) {
				if _, hasDefaultScanner := repl.scanner.(*bufio.Scanner); !hasDefaultScanner {
					t.Errorf("repl.Scanner is not the default bufio.NewScanner(os.Stdin) instance")
				}
			})

			t.Run("should have the default prompt", func(t *testing.T) {
				if repl.prompt != defaultPrompt {
					t.Errorf("repl.Prompt is not the default prompt")
				}
			})
		})
	})

	t.Run("REPL creation error handling", func(t *testing.T) {
		t.Run("should panic if command executor is not set", func(t *testing.T) {
			defer func() {
				if recover := recover(); recover == nil {
					t.Errorf("Expected NewREPL to panic, but it didn't")
				}
			}()

			NewREPL()
		})
	})
}

// @SECTION: MOCKS

func commandExecutorMock(args string) error {
	return nil
}
