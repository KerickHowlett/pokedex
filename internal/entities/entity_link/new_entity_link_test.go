package entity_link

import (
	"testing"

	f "test_tools/fixtures"
)

func TestNewEntityLink(t *testing.T) {
	t.Run("should create a new entity link.", func(t *testing.T) {
		if link := NewEntityLink(); link == nil {
			t.Error("expected link to not be nil.")
		}
	})

	t.Run("should set the name of the entity link.", func(t *testing.T) {
		if link := NewEntityLink(WithName(f.PokemonName)); link.Name != f.PokemonName {
			t.Errorf("expected name to be Pikachu; got %s", link.Name)
		}
	})
}
