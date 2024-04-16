package help_command

import (
	"reflect"
	"test_tools/utils"
	"testing"
)

func TestHelpCommand_Execute(t *testing.T) {
	t.Run("should panic.", func(t *testing.T) {
		command := NewHelpCommand()
		utils.ExpectPanic(t, command.Execute)
	})
}

func TestHelpCommand_GetArgs(t *testing.T) {
	t.Run("should return an empty slice.", func(t *testing.T) {
		if command := NewHelpCommand(); reflect.DeepEqual(command.GetArgs(), []string{}) == false {
			t.Error("Expected GetArgs to return an empty slice, but it didn't.")
		}
	})
}

func TestHelpCommand_GetDescription(t *testing.T) {
	t.Run("should return the description of the HelpCommand.", func(t *testing.T) {
		if command := NewHelpCommand(); command.GetDescription() != "Show help for the Pokemon CLI commands." {
			t.Errorf("Expected GetDescription to return 'Show help for the Pokemon CLI commands.', but got '%s'", command.GetDescription())
		}
	})
}

func TestHelpCommand_GetName(t *testing.T) {
	t.Run("should return the name of the HelpCommand.", func(t *testing.T) {
		if command := NewHelpCommand(); command.GetName() != "help" {
			t.Errorf("Expected GetName to return 'help', but got '%s'", command.GetName())
		}
	})
}

func TestHelpCommand_PrintHelp(t *testing.T) {
	stdout := utils.NewPrintStorage()
	t.Run("should call the PrintHelp function.", func(t *testing.T) {
		command := NewHelpCommand()
		if output := stdout.Capture(command.PrintHelp); output == "" {
			t.Error("Expected PrintHelp to be called, but it wasn't.")
		}
	})
}

func TestHelpCommand_SetArgs(t *testing.T) {
	t.Run("should do nothing.", func(t *testing.T) {
		command := NewHelpCommand()
		command.SetArgs([]string{"arg1", "arg2"})
		if reflect.DeepEqual(command.GetArgs(), []string{}) == false {
			t.Error("Expected SetArgs to do nothing, but it didn't.")
		}
	})
}
