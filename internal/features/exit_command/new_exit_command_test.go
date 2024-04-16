package exit_command

import "testing"

func TestNewExitCommand(t *testing.T) {
	t.Run("should return a new instance of ExitCommand.", func(t *testing.T) {
		if command := NewExitCommand(); command == nil {
			t.Errorf("the returned command was nil")
		}
	})
}
