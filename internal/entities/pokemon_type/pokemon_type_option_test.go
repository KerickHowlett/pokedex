package pokemon_type

import (
	el "entity_link"
	"testing"
)

func TestWithSlot(t *testing.T) {
	runWithSlotTest := func() (pokemonType *PokemonType, slot int) {
		slot = 1
		pokemonType = &PokemonType{}
		WithSlot(slot)(pokemonType)
		return pokemonType, slot
	}

	t.Run("should set the 'slot' field.", func(t *testing.T) {
		if pokemonType, slot := runWithSlotTest(); pokemonType.Slot != slot {
			t.Errorf("WithSlot() = %v; wanted %v", pokemonType.Slot, slot)
		}
	})
}

func TestWithType(t *testing.T) {
	runWithTypeTest := func() (pokemonType *PokemonType, pokemonTypeLink *el.EntityLink) {
		pokemonTypeLink = el.NewEntityLink()
		pokemonType = &PokemonType{}
		WithType(pokemonTypeLink)(pokemonType)
		return pokemonType, pokemonTypeLink
	}

	t.Run("should set the 'type' field.", func(t *testing.T) {
		if pokemonType, pokemonTypeLink := runWithTypeTest(); pokemonType.Type != pokemonTypeLink {
			t.Errorf("WithType() = %v; wanted %v", pokemonType.Type, pokemonTypeLink)
		}
	})
}
