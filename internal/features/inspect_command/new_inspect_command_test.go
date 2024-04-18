package inspect_command

import "testing"

func TestNewInspectCommand(t *testing.T) {
	t.Run("Given no provided options", func(t *testing.T) {
		t.Run("When a new InspectCommand struct is instantiated", func(t *testing.T) {
			t.Run("Then the 'description' field should be set with it's default value", func(t *testing.T) {
				t.Parallel()
				if command := NewInspectCommand(); command.description != "Take a closer look at one of your Pokemon." {
					t.Errorf("expected description to be 'Take a closer look at one of your Pokemon.', got %s", command.description)
				}
			})

			t.Run("And the 'name' field should be set with it's default value", func(t *testing.T) {
				t.Parallel()
				if command := NewInspectCommand(); command.name != "inspect" {
					t.Errorf("expected name to be 'inspect', got %s", command.name)
				}
			})

			t.Run("And the 'pc' field should be set with a non-nil value", func(t *testing.T) {
				t.Parallel()
				// @NOTE: There's no way to test if it's the same function given how its implemented.
				if command := NewInspectCommand(); command.pc == nil {
					t.Error("expected pc to be a non-nil value")
				}
			})
		})
	})
}
