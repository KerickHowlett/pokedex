package toolchain

import (
	"fmt"
	"testing"

	m "internal/tests/mocks/command"
)

func TestWithCommand(t *testing.T) {
	fmt.Println("Should add a command to the toolchain.")
	command := m.NewMockCommand()
	toolchain := NewToolchain(WithCommand(command))

	if _, exists := toolchain.SelectCommand(command.GetName()); !exists {
		t.Errorf("Expected command to exist in toolchain")
	}
}
