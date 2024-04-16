package exit_command

import "testing"

func TestNewExitCommand(t *testing.T) {
	t.Run("should return an instance of ExitCommand.", func(t *testing.T) {
		t.Parallel()
		if command := NewExitCommand(); command == nil {
			t.Error("Expected command to be an instance of ExitCommand, but got nil.")
		}
	})

	t.Run("should be instantiated with the default value for it's 'description' field.", func(t *testing.T) {
		t.Parallel()
		if command := NewExitCommand(); command.description != "Exit the Pokemon CLI application." {
			t.Errorf("Expected description to be 'Exit the Pokemon CLI application.', but got '%s'", command.description)
		}
	})

	t.Run("should be instantiated with the default value for it's 'name' field.", func(t *testing.T) {
		t.Parallel()
		if command := NewExitCommand(); command.name != "exit" {
			t.Errorf("Expected name to be 'exit', but got '%s'", command.name)
		}
	})
}
