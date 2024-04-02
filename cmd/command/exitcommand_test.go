package command

import (
	"os"
	"os/exec"
	"testing"
)

// @SECTION: Unit Test Cases

func TestExitCommand(t *testing.T) {

	// This function configures the TestBed environment,
	// so the test case(s) doesn't exit prematurely.
	setup := func() error {
		if os.Getenv("BE_CRASHER") == "1" {
			command := ExitCommand{}

			return command.Execute()
		}

		cmd := exec.Command(os.Args[0], "-test.run=TestExitProcess")
		cmd.Env = append(os.Environ(), "BE_CRASHER=1")

		return cmd.Run()
	}

	t.Run("should exit any ongoing process gracefully with status 0.", func(t *testing.T) {
		err := setup()

		if _, ok := err.(*exec.ExitError); !ok {
			return
		}

		t.Fatalf("Process exited with error %v, expected status 0", err)
	})
}
