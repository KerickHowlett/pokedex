package inspect_command

import (
	el "entity_link"
	ps "pokemon_stat"
	pt "pokemon_type"
	"testing"

	p "pokemon"
	f "test_tools/fixtures"
	"test_tools/utils"
)

func Test_printBasicInfoOf(t *testing.T) {
	t.Run("should print the basic information of a pokemon", func(t *testing.T) {
		expectStdoutFrom(printBasicInfoOf, "Name: Pikachu\nHeight: 4\nWeight: 60\n", t)
	})
}

func Test_printTypesOf(t *testing.T) {
	t.Run("should print the subject pokemon's given types", func(t *testing.T) {
		expectStdoutFrom(printTypesOf, "Types:\n - Electric:\n", t)
	})
}

func Test_printStatsOf(t *testing.T) {
	t.Run("should print the subject pokemon's given stats", func(t *testing.T) {
		expectStdoutFrom(printStatsOf, "Stats:\n - hp: 35\n", t)
	})
}

func expectStdoutFrom(fnc func(pokemon *p.Pokemon), expected string, t *testing.T) {
	pokemon := &p.Pokemon{
		Name:   f.PokemonName,
		Height: 4,
		Weight: 60,
		Types: []*pt.PokemonType{{
			Type: &el.EntityLink{Name: "Electric"},
		}},
		Stats: []*ps.PokemonStat{{
			BaseStat: 35,
			Stat:     &el.EntityLink{Name: "hp"},
		}},
	}
	stdout := utils.NewPrintStorage()
	if actual := stdout.Capture(func() { fnc(pokemon) }); actual != expected {
		t.Errorf("Expected '%s' to be printed out, only to have this printed out: %s", expected, actual)
	}
}
