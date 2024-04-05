package utils

import (
	"os"
	"os/exec"
	"testing"
)

func TestExitApplication(t *testing.T) {

	// This function configures the TestBed environment,
	// so the test case(s) doesn't exit prematurely.
	setup := func() error {
		if os.Getenv("BE_CRASHER") == "1" {
			return ExitApplication(OK)
		}

		cmd := exec.Command(os.Args[0], "-test.run=TestExitProcess")
		cmd.Env = append(os.Environ(), "BE_CRASHER=1")

		return cmd.Run()
	}

	t.Run("should exit application gracefully with OK status.", func(t *testing.T) {
		err := setup()
		if _, ok := err.(*exec.ExitError); ok {
			t.Fatalf("Process exited with error %v, expected exit status of %v", err, OK)
		}
	})
}
