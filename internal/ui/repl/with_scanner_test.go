package repl

import (
	"testing"

	mc "test_tools/mocks/command_executor"
	ms "test_tools/mocks/scanner"
	u "test_tools/utils"
)

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
		u.ExpectSameEntity(t, repl.scanner, scanner, "Scanner")
	})
}
