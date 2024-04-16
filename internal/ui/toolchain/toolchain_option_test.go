package toolchain

import (
	"testing"

	m "test_tools/mocks/command"
	"test_tools/utils"
)

func TestWithCommand(t *testing.T) {
	command := m.NewMockCommand()

	t.Run("should add the command to the toolchain.", func(t *testing.T) {
		toolchain := NewToolchain(WithCommand(command))
		commandName := command.GetName()

		if _, commandExists := toolchain.SelectCommand(commandName); !commandExists {
			t.Errorf("Expected command to exist in toolchain")
		}
	})

	t.Run("should panic to prevent multiple commands with the same name from being added to the toolchain.", func(t *testing.T) {
		utils.ExpectPanic(t, func() {
			NewToolchain(WithCommand(command), WithCommand(command))
		})
	})
}
