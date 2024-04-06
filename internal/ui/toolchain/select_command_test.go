package toolchain

import (
	"testing"

	m "test_tools/mocks/command"
)

func TestToolchain_SelectCommand(t *testing.T) {
	const defaultCommandName string = "mock1"

	setup := func(targetMockCommandName string) (toolchain *Toolchain, command1 *m.MockCommand, command2 *m.MockCommand) {
		command1 = m.NewMockCommand(m.WithName(targetMockCommandName))
		command2 = m.NewMockCommand(m.WithName("mock2"))

		toolchain = NewToolchain(
			WithCommand(command1),
			WithCommand(command2),
		)

		return toolchain, command1, command2
	}

	t.Run("should select the correct command when queried by name.", func(t *testing.T) {
		toolchain, command1, _ := setup(defaultCommandName)
		if foundCommand, _ := toolchain.SelectCommand(defaultCommandName); (*foundCommand).GetName() != command1.GetName() {
			t.Errorf("Expected foundCommand to be command1, but got %v", foundCommand)
		}
	})

	t.Run("should return a true boolean value if the command exists.", func(t *testing.T) {
		toolchain, _, _ := setup(defaultCommandName)
		if _, commandExists := toolchain.SelectCommand(defaultCommandName); !commandExists {
			t.Errorf("Expected commandExists to be true, but got false")
		}
	})

	t.Run("should return a false boolean value if the command does not exist.", func(t *testing.T) {
		toolchain, _, _ := setup(defaultCommandName)
		if _, commandExists := toolchain.SelectCommand("nonexistent"); commandExists {
			t.Errorf("Expected commandExists to be false, but got true")
		}
	})
}
