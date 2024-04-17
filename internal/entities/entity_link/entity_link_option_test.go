package entity_link

import (
	f "test_tools/fixtures"
	"testing"
)

func TestWithName(t *testing.T) {
	runWithNameTest := func() (entityLink *EntityLink) {
		entityLink = &EntityLink{}
		WithName(f.PokemonName)(entityLink)
		return entityLink
	}

	t.Run("should set name field", func(t *testing.T) {
		if entityLink := runWithNameTest(); entityLink.Name != f.PokemonName {
			t.Errorf("expected name to be %s; got %s", f.PokemonName, entityLink.Name)
		}
	})
}

func TestWithURL(t *testing.T) {
	runWithURLTest := func() (entityLink *EntityLink) {
		entityLink = &EntityLink{}
		WithURL(f.APIEndpoint)(entityLink)
		return entityLink
	}

	t.Run("should set URL field", func(t *testing.T) {
		if entityLink := runWithURLTest(); entityLink.URL != f.APIEndpoint {
			t.Errorf("expected URL to be %s; got %s", f.APIEndpoint, entityLink.URL)
		}
	})
}
