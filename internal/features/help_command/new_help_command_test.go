package help_command

import "testing"

func TestNewExitCommand(t *testing.T) {
	t.Run("should return an instance of ExitCommand.", func(t *testing.T) {
		t.Parallel()
		if command := NewHelpCommand(); command == nil {
			t.Error("Expected command to be an instance of ExitCommand, but got nil.")
		}
	})

	t.Run("should be instantiated with the default value for it's 'description' field.", func(t *testing.T) {
		t.Parallel()
		if command := NewHelpCommand(); command.description != "Show help for the Pokemon CLI commands." {
			t.Errorf("Expected description to be 'Show help for the Pokemon CLI commands.', but got '%s'", command.description)
		}
	})

	t.Run("should be instantiated with the default value for it's 'name' field.", func(t *testing.T) {
		t.Parallel()
		if command := NewHelpCommand(); command.name != "help" {
			t.Errorf("Expected name to be 'help', but got '%s'", command.name)
		}
	})
}
