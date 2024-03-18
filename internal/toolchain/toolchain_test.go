package toolchain

import (
	"fmt"
	"testing"

	m "github.com/KerickHowlett/pokedexcli/tests/mocks"
)

func TestNewToolchain(t *testing.T) {
	fmt.Println("Should create a new toolchain with at least one command.")
	toolchain := NewToolchain(
		WithCommand(m.NewMockCommand()),
	)

	if len(*toolchain.commands) == 0 {
		t.Errorf("Expected Toolchain to have at least one command, but got 0")
	}
}

func TestNewToolchain_Panic(t *testing.T) {
	fmt.Println("Should panic if the toolchain has no commands.")
	defer func() {
		if recover := recover(); recover == nil {
			t.Errorf("Expected NewToolchain to panic, but it didn't")
		}
	}()

	NewToolchain()
}
