package commands

import (
	"testing"
)

// @TODO: Replace with something more dynamic.

func TestExitCommandIsDefined(t *testing.T) {
	commands := GetCommands()

	assertCommandIsDefined(t, commands, "exit")
}

func TestHelpCommandIsDefined(t *testing.T) {
	commands := GetCommands()

	assertCommandIsDefined(t, commands, "help")
}

func assertCommandIsDefined(t *testing.T, commands Commands, commandName string) {
	if _, ok := commands[commandName]; !ok {
		t.Errorf("Expected command '%s' to be present, but it is not", commandName)
	}
}
