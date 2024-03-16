package command

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

// @SECTION: Unit Test Cases

func TestExitProcess(t *testing.T) {
	fmt.Println("Should exit any ongoing process gracefully with status 0.")

	err := setupExitCommandTestBed()

	if _, ok := err.(*exec.ExitError); !ok {
		return
	}

	t.Fatalf("Process exited with error %v, expected status 0", err)
}

// @SECTION: Helper Functions

// This function configures the TestBed environment,
// so the test case(s) doesn't exit prematurely.
func setupExitCommandTestBed() error {
	if os.Getenv("BE_CRASHER") == "1" {
		command := ExitCommand{}

		return command.Execute()
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestExitProcess")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")

	return cmd.Run()
}
