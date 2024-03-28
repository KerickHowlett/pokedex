package toolchain

import (
	"fmt"
	"testing"

	c "github.com/KerickHowlett/pokedexcli/cmd/command"
)

func TestWithCommand(t *testing.T) {
	fmt.Println("Should add a command to the toolchain.")
	command := c.NewHelpCommand()
	toolchain := NewToolchain(WithCommand(command))

	if _, exists := toolchain.SelectCommand(command.GetName()); !exists {
		t.Errorf("Expected command to exist in toolchain")
	}
}
