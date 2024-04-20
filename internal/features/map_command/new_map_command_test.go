package map_command

import (
	"testing"

	pd "map_command/pagination_direction"
	f "test_tools/fixtures"
)

func TestNewMapCommand(t *testing.T) {
	setup := func(paginationDirection string) *MapCommand {
		return NewMapCommand(WithPaginationDirection(paginationDirection))
	}

	t.Run("Given the function was provided with either 'pagination direction' option", func(t *testing.T) {
		t.Run("should have the default value for the cacheTTL.", func(t *testing.T) {
			// @NOTE: There's no way to test if it's the same function given how its implemented.
			if command := setup(pd.Next); command.ec == nil {
				t.Error("expected ec to be defined but instead got nil")
			}
		})

		t.Run("should have the default value for the fetchLocation.", func(t *testing.T) {
			// @NOTE: There's no way to test if it's the same function given how its implemented.
			if command := setup(pd.Next); command.fetchLocations == nil {
				t.Error("expected fetchLocations to be a function possessing the correct call signature but instead got nil")
			}
		})

		t.Run("should have the default value for the listMarker.", func(t *testing.T) {
			if command := setup(pd.Next); command.listMarker != "  -" {
				t.Errorf("expected listMarker to be '  -', got %v", command.listMarker)
			}
		})

		t.Run("should have the default value for the listTitle.", func(t *testing.T) {
			if command := setup(pd.Next); command.listTitle != "Pokemon Maps:" {
				t.Errorf("expected listTitle to be 'Pokemon Maps:', got %v", command.listTitle)
			}
		})

		t.Run("should have the default value for the noMapsFoundErrorMessage.", func(t *testing.T) {
			if command := setup(pd.Next); command.noMapsFoundErrorMessage != "no maps were found" {
				t.Errorf("expected noMapsFoundErrorMessage to be 'no maps were found', got %v", command.noMapsFoundErrorMessage)
			}
		})

		t.Run("should have the default value for the noMoreMapsMessage.", func(t *testing.T) {
			if command := setup(pd.Next); command.noMoreMapsMessage != "There are no more maps." {
				t.Errorf("expected noMoreMapsMessage to be 'There are no more maps.', got %v", command.noMoreMapsMessage)
			}
		})

		t.Run("should have the default value for the state.", func(t *testing.T) {
			// @NOTE: There's no way to test if it's the same function given how its implemented.
			if command := setup(pd.Next); command.state == nil {
				t.Error("expected state to be defined but instead got nil")
			}
		})
	})

	t.Run("Given the function was provided with the 'pagination direction' option of 'next'", func(t *testing.T) {
		t.Parallel()
		t.Run("should have the default value for the commandName.", func(t *testing.T) {
			if command := setup(pd.Next); command.name != "map" {
				t.Errorf("expected commandName to be 'map', got %v", command.name)
			}
		})

		t.Run("should have the default value for the commandDescription.", func(t *testing.T) {
			if command := setup(pd.Next); command.description != "Show the next 20 Pokemon world map locations." {
				t.Errorf("expected commandDescription to be 'Show the next 20 Pokemon world map locations.', got %v", command.description)
			}
		})

		t.Run("should have the default value for the noMoreMapsMessage.", func(t *testing.T) {
			if command := setup(pd.Next); command.noMoreMapsMessage != "There are no more maps." {
				t.Errorf("expected noMoreMapsMessage to be 'There are no more maps.', got %v", command.noMoreMapsMessage)
			}
		})
	})

	t.Run("Given the function was provided with the 'pagination direction' option of 'previous'", func(t *testing.T) {
		t.Parallel()
		t.Run("should have the default value for the commandName.", func(t *testing.T) {
			if command := setup(pd.Previous); command.name != "mapb" {
				t.Errorf("expected commandName to be 'mapb', got %v", command.name)
			}
		})

		t.Run("should have the default value for the commandDescription.", func(t *testing.T) {
			if command := setup(pd.Previous); command.description != "Show the previous 20 Pokemon world map locations." {
				t.Errorf("expected commandDescription to be 'Show the previous 20 Pokemon world map locations.', got %v", command.description)
			}
		})

		t.Run("should have the default value for the noMoreMapsMessage.", func(t *testing.T) {
			if command := setup(pd.Previous); command.noMoreMapsMessage != "You're already at the beginning of the maps list." {
				t.Errorf("expected noMoreMapsMessage to be 'You're already at the beginning of the maps list.', got %v", command.noMoreMapsMessage)
			}
		})
	})

	t.Run("Given the function was provided with a 'pagination direction' option of 'invalid'", func(t *testing.T) {
		t.Parallel()
		t.Run("should panic with the message 'Invalid pagination direction'", func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Error("Expected panic, but no panic occurred")
				}
			}()

			setup(f.Invalid)
		})
	})

	t.Run("Given the function was provided with no options", func(t *testing.T) {
		t.Parallel()
		t.Run("should panic if fetchLocations is not set", func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Error("Expected panic, but no panic occurred")
				}
			}()

			NewMapCommand(func(c *MapCommand) {
				c.fetchLocations = nil
			})
		})
	})
}

func TestSetDefaultCommandName(t *testing.T) {
	t.Run("should return 'map' for pagination direction 'next'", func(t *testing.T) {
		t.Parallel()
		expected := "map"
		if actual := setDefaultCommandName(pd.Next); actual != expected {
			t.Errorf("Expected %s, but got %s", expected, actual)
		}
	})

	t.Run("should return 'mapb' for pagination direction 'previous'", func(t *testing.T) {
		t.Parallel()
		expected := "mapb"
		if actual := setDefaultCommandName(pd.Previous); actual != expected {
			t.Errorf("Expected %s, but got %s", expected, actual)
		}
	})

	t.Run("should panic for invalid pagination direction", func(t *testing.T) {
		t.Parallel()
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic, but no panic occurred")
			}
		}()
		setDefaultCommandName(f.Invalid)
	})
}

func TestSetDefaultNoMoreMapsMessage(t *testing.T) {
	t.Run("should return 'There are no more maps.' for pagination direction 'next'", func(t *testing.T) {
		t.Parallel()
		expected := "There are no more maps."
		if actual := setDefaultNoMoreMapsMessage(pd.Next); actual != expected {
			t.Errorf("Expected %s, but got %s", expected, actual)
		}
	})

	t.Run("should return 'You're already at the beginning of the maps list.' for pagination direction 'previous'", func(t *testing.T) {
		t.Parallel()
		expected := "You're already at the beginning of the maps list."
		if actual := setDefaultNoMoreMapsMessage(pd.Previous); actual != expected {
			t.Errorf("Expected %s, but got %s", expected, actual)
		}
	})

	t.Run("should panic for invalid pagination direction", func(t *testing.T) {
		t.Parallel()
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic, but no panic occurred")
			}
		}()
		setDefaultNoMoreMapsMessage(f.Invalid)
	})
}
