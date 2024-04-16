package exit_command

import (
	"os"
	"os/exec"
	"test_tools/utils"
	"testing"
)

func TestGetDescription(t *testing.T) {
	t.Run("should return the description of the ExitCommand.", func(t *testing.T) {
		if command := NewExitCommand(); command.GetDescription() != "Exit the Pokemon CLI application." {
			t.Errorf("Expected GetDescription to return 'Exit the Pokemon CLI application.', but got '%s'", command.GetDescription())
		}
	})
}

func TestGetName(t *testing.T) {
	t.Run("should return the name of the ExitCommand.", func(t *testing.T) {
		if command := NewExitCommand(); command.GetName() != "exit" {
			t.Errorf("Expected GetName to return 'exit', but got '%s'", command.GetName())
		}
	})
}

func TestGetArgs(t *testing.T) {
	t.Run("should return an empty slice.", func(t *testing.T) {
		if command := NewExitCommand(); len(command.GetArgs()) != 0 {
			t.Error("Expected GetArgs to return an empty slice, but it didn't.")
		}
	})
}

func TestExecute(t *testing.T) {
	t.Run("should terminate the program with an exit status of 0.", func(t *testing.T) {
		if os.Getenv("BE_CRASHER") == "1" {
			command := NewExitCommand()
			err := command.Execute()
			if err != nil {
				t.Errorf("Expected Execute to return nil, but got %v", err)
			}
			return
		}

		cmd := exec.Command(os.Args[0], "-test.run=TestExitCommand")
		cmd.Env = append(os.Environ(), "BE_CRASHER=1")
		err := cmd.Run()
		if e, ok := err.(*exec.ExitError); ok && !e.Success() {
			t.Errorf("Process ran with err %v, want exit status 0", e)
		}
	})
}

func TestPrintHelp(t *testing.T) {
	stdout := utils.NewPrintStorage()
	t.Run("should call the PrintHelp function.", func(t *testing.T) {
		command := NewExitCommand()
		if output := stdout.Capture(command.PrintHelp); output == "" {
			t.Error("Expected PrintHelp to be called, but it wasn't.")
		}
	})
}

func TestSetArgs(t *testing.T) {
	t.Run("should do nothing.", func(t *testing.T) {
		command := NewExitCommand()
		command.SetArgs([]string{"arg1", "arg2"})
		if len(command.GetArgs()) != 0 {
			t.Error("Expected SetArgs to do nothing, but it didn't.")
		}
	})
}
