package commands

import (
	"os"
	"os/exec"
	"testing"
)

func setup() error {
	if os.Getenv("BE_CRASHER") == "1" {
		ExitREPL()
		return nil
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestExitREPL")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()

	return err
}

func TestExitREPL(t *testing.T) {
	err := setup()

	if _, ok := err.(*exec.ExitError); !ok {
		return
	}

	t.Fatalf("Process ran with error %v, expected status 0", err)
}
