package commands

import (
	"fmt"
	"testing"
)

// @SECTION: Unite Test Cases

func TestRunCommandWithEmptyCommand(t *testing.T) {
	fmt.Println("Given an empty command selection, then RunCommand() should not call the command executioner.")

	commands, isFunctionCalled := setupRunCommandTestBed()

	if response := Run("", commands); response != nil {
		t.Errorf("Expected no errors, but received: %v", response)
	}

	if *isFunctionCalled {
		t.Error("Expected the command executor to not be called.")
	}
}

func TestRunCommandWithValidCommand(t *testing.T) {
	fmt.Println("Given a valid command selection, then RunCommand() should call the command executioner.")

	commands, isFunctionCalled := setupRunCommandTestBed()

	Run("test", commands)

	if !*isFunctionCalled {
		t.Error("Expected Execute() to be called")
	}
}

func TestRunCommandWithNoErrorsReturned(t *testing.T) {
	fmt.Println("Given a valid command selection, then RunCommand() should return no errors.")

	commands, _ := setupRunCommandTestBed()

	if err := Run("test", commands); err != nil {
		t.Errorf("Expected no errors but got: %v", err)
	}
}

func TestRunCommandWithInvalidCommand(t *testing.T) {
	fmt.Println("Given an invalid command selection, then RunCommand() should return an error.")

	commands, _ := setupRunCommandTestBed()
	expectedError := fmt.Errorf("Command not found: foobar")

	if err := Run("foobar", commands); err.Error() != expectedError.Error() {
		t.Errorf("Expected error: %v, but got: %v", expectedError, err)
	}
}

// @SECTION: Helper Functions

func setupRunCommandTestBed() (commands Commands, isFunctionCalled *bool) {
	isFunctionCalled = new(bool)
	*isFunctionCalled = false

	commands = Commands{
		"test": {
			Name:        "test",
			Description: "This is a fake command.",
			Execute: func() error {
				*isFunctionCalled = true
				return nil
			},
		},
	}

	return commands, isFunctionCalled
}
