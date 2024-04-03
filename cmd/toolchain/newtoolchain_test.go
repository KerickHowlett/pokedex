package toolchain

import (
	"testing"

	m "testtools/mocks/command"
)

func TestNewToolchain(t *testing.T) {
	t.Run("should create a new toolchain successfully with one command requirement.", func(t *testing.T) {
		if toolchain := NewToolchain(WithCommand(m.NewMockCommand())); len(*toolchain.commands) == 0 {
			t.Errorf("Expected Toolchain to have at least one command, but got 0")
		}
	})

	t.Run("should panic if the toolchain was given no commands for initialization.", func(t *testing.T) {
		defer func() {
			if recover := recover(); recover == nil {
				t.Errorf("Expected NewToolchain to panic, but it didn't")
			}
		}()
		NewToolchain()
	})
}
