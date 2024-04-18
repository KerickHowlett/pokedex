package help_command

import (
	"reflect"
	"testing"

	"test_tools/utils"
)

func TestHelpCommand_Execute(t *testing.T) {
	t.Run("should panic.", func(t *testing.T) {
		command := &HelpCommand{}
		utils.ExpectPanic(t, command.Execute)
	})
}

func TestHelpCommand_GetArgs(t *testing.T) {
	t.Run("should return an empty slice.", func(t *testing.T) {
		command := &HelpCommand{}
		if reflect.DeepEqual(command.GetArgs(), []string{}) == false {
			t.Error("Expected GetArgs to return an empty slice, but it didn't.")
		}
	})
}

func TestHelpCommand_GetDescription(t *testing.T) {
	setup := func() (actual string, expected string) {
		expected = "Show help for the Pokemon CLI commands."
		command := &HelpCommand{description: expected}
		actual = command.GetDescription()
		return actual, expected
	}
	t.Run("should return the description of the HelpCommand.", func(t *testing.T) {
		if actual, expected := setup(); actual != expected {
			t.Errorf("Expected GetDescription to return '%s', but got '%s'", expected, actual)
		}
	})
}

func TestHelpCommand_GetName(t *testing.T) {
	setup := func() (actual string, expected string) {
		expected = "help"
		command := &HelpCommand{name: expected}
		actual = command.GetName()
		return actual, expected
	}
	t.Run("should return the name of the HelpCommand.", func(t *testing.T) {
		if actual, expected := setup(); actual != expected {
			t.Errorf("Expected GetName to return '%s', but got '%s'", expected, actual)
		}
	})
}

func TestHelpCommand_PrintHelp(t *testing.T) {
	stdout := utils.NewPrintStorage()
	setup := func() string {
		command := &HelpCommand{description: "foobar", name: "foobar"}
		command.PrintHelp()
		return stdout.Capture(command.PrintHelp)

	}
	t.Run("should call the PrintHelp function.", func(t *testing.T) {
		if output := setup(); output == "" {
			t.Error("Expected PrintHelp to be called, but it wasn't.")
		}
	})
}

func TestHelpCommand_SetArgs(t *testing.T) {
	setup := func() []string {
		command := &HelpCommand{}
		command.SetArgs([]string{"foo", "bar"})
		actual := command.GetArgs()
		return actual
	}
	t.Run("should do nothing.", func(t *testing.T) {
		if actual := setup(); !reflect.DeepEqual(actual, []string{}) {
			t.Error("Expected SetArgs to do nothing, but it didn't.")
		}
	})
}
