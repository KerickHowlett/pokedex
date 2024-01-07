package commands

import (
	"fmt"
	"testing"
)

func mockGetCommands() (oldGetCommands func() Commands, executeCalled bool) {
	executeCalled = false

	oldGetCommands = GetCommands
	GetCommands = func() Commands {
		return Commands{
			"exit": {
				Name:        "exit",
				Description: "Exits the application",
				Execute: func() error {
					executeCalled = true
					return nil
				},
			},
			"help": {
				Name:        "help",
				Description: "Displays the help menu",
				Execute: func() error {
					executeCalled = true
					return nil
				},
			},
		}
	}

	return oldGetCommands, executeCalled
}

func afterEach(oldGetCommands func() Commands) {
	GetCommands = oldGetCommands
}

func TestRunCommandWithEmptyCommand(t *testing.T) {
	err := RunCommand("")

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestRunCommandWithValidCommand(t *testing.T) {
	oldGetCommands, executeCalled := mockGetCommands()
	defer afterEach(oldGetCommands)

	RunCommand("exit")

	if !executeCalled {
		t.Error("Expected Execute() to be called")
	}
}

func TestRunCommandWithInvalidCommand(t *testing.T) {
	oldGetCommands, _ := mockGetCommands()
	defer afterEach(oldGetCommands)

	expectedError := fmt.Errorf("Command not found: unknown")

	err := RunCommand("unknown")

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error: %v, but got: %v", expectedError, err)
	}
}
