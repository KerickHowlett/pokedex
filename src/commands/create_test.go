package commands

import (
	"fmt"
	"testing"
)

// * This is a fake Commands fixture to be used
// * for testing throughout the 'commands' package.

func TestExitCommandIsDefined(t *testing.T) {
	commands := Create()

	assertCommandIsDefined(t, commands, "exit")
}

func TestHelpCommandIsDefined(t *testing.T) {
	commands := Create()

	assertCommandIsDefined(t, commands, "help")
}

func assertCommandIsDefined(t *testing.T, commands Commands, commandName string) {
	fmt.Printf("Create() should possess the '%s' command.", commandName)

	if _, ok := commands[commandName]; !ok {
		t.Errorf("Expected command '%s' to be present, but it is not", commandName)
	}
}
