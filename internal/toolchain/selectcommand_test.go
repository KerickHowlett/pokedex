package toolchain

import (
	"fmt"
	"testing"

	m "github.com/KerickHowlett/pokedexcli/tests/mocks"
)

const defaultCommandName string = "mock1"

func TestToolchain_SelectCommand(t *testing.T) {
	fmt.Println("Should select the correct command when queried by name.")
	toolchain, command1, _ := setupSelectCommandTests(defaultCommandName)

	foundCommand, _ := toolchain.SelectCommand(defaultCommandName)
	if (*foundCommand).GetName() != command1.GetName() {
		t.Errorf("Expected foundCommand to be command1, but got %v", foundCommand)
	}
}

func TestToolchain_SelectCommand_CommandExists(t *testing.T) {
	fmt.Println("Should return a true boolean value if the command exists.")
	toolchain, _, _ := setupSelectCommandTests(defaultCommandName)

	if _, commandExists := toolchain.SelectCommand(defaultCommandName); !commandExists {
		t.Errorf("Expected commandExists to be true, but got false")
	}
}

func TestToolchain_SelectCommand_CommandDoesNotExist(t *testing.T) {
	fmt.Println("Should return a false boolean value if the command does not exist.")
	toolchain, _, _ := setupSelectCommandTests(defaultCommandName)

	if _, commandExists := toolchain.SelectCommand("nonexistent"); commandExists {
		t.Errorf("Expected commandExists to be false, but got true")
	}
}

func setupSelectCommandTests(targetMockCommandName string) (*Toolchain, *m.MockCommand, *m.MockCommand) {
	command1 := m.NewMockCommand()
	command1.SetName(targetMockCommandName)

	command2 := m.NewMockCommand()
	command2.SetName("mock2")

	toolchain := NewToolchain(
		WithCommand(command1),
		WithCommand(command2),
	)

	return toolchain, command1, command2
}
