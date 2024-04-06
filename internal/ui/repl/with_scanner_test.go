package repl

import (
	"testing"

	m "test_tools/mocks/scanner"
	u "test_tools/utils"
)

func TestWithScanner(t *testing.T) {
	setup := func() (repl *REPL, scanner *m.MockScanner) {
		scanner = m.NewMockScanner()
		repl = NewREPL(
			WithCommandExecutor(commandExecutorMock),
			WithScanner(scanner),
		)

		return repl, scanner
	}

	t.Run("should set the Scanner field of the REPL struct", func(t *testing.T) {
		repl, scanner := setup()
		u.ExpectSameEntity(t, repl.scanner, scanner, "Scanner")
	})
}
