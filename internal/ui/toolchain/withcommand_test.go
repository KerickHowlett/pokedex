package toolchain

import (
	"testing"

	m "testtools/mocks/command"
)

func TestWithCommand(t *testing.T) {
	setup := func() (command *m.MockCommand, toolchain *Toolchain) {
		command = m.NewMockCommand()
		toolchain = NewToolchain(WithCommand(command))

		return command, toolchain
	}

	t.Run("should add the command to the toolchain.", func(t *testing.T) {
		command, toolchain := setup()
		commandName := command.GetName()

		if _, commandExists := toolchain.SelectCommand(commandName); !commandExists {
			t.Errorf("Expected command to exist in toolchain")
		}
	})
}
