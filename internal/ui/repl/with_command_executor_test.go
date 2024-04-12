package repl

import (
	"testing"

	m "test_tools/mocks/command_executor"
	"test_tools/utils"
)

func TestWithCommandExecutor(t *testing.T) {
	setup := func() func(command string, args []string) error { // Modify the function signature
		repl := NewREPL(WithCommandExecutor(m.MockedCommandExecutor))
		return repl.execute
	}

	t.Run("should set the Fields field of the Sanitizer struct", func(t *testing.T) {
		execute := setup()
		utils.ExpectSameEntity(t, execute, m.MockedCommandExecutor, "execute")
	})
}
