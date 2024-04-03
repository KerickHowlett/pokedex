package repl

import (
	"testing"

	m "testtools/mocks/commandexecutor"
	"testtools/utils"
)

func TestWithCommandExecutor(t *testing.T) {
	setup := func() func(string) error {
		repl := NewREPL(WithCommandExecutor(m.MockedCommandExecutor))
		return repl.execute
	}

	t.Run("should set the Fields field of the Sanitizer struct", func(t *testing.T) {
		execute := setup()
		utils.ExpectSameEntity(t, execute, m.MockedCommandExecutor, "execute")
	})
}
