package help_command

import "testing"

func TestNewHelpCommand(t *testing.T) {
	t.Run("should return a new instance of ExitCommand.", func(t *testing.T) {
		if command := NewHelpCommand(); command == nil {
			t.Errorf("the returned command was nil")
		}
	})
}
