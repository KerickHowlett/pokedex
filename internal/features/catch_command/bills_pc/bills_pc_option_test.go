package bills_pc

import (
	p "pokemon"
	f "test_tools/fixtures"
	"testing"
)

func TestWithCaughtPokemon(t *testing.T) {
	t.Run("should set the caught Pokemon in the BillsPC state.", func(t *testing.T) {
		pokemon := p.NewPokemon(p.WithName(f.PokemonName))
		billsPC := NewBillsPC(WithCaughtPokemon(pokemon))
		if _, found := (*billsPC.caughtPokemon)[f.PokemonName]; !found {
			t.Errorf("Expected to find the caught Pokemon in the BillsPC state, but it was not found.")
		}
	})
}

func TestWithWithdrawError(t *testing.T) {
	t.Run("should set the withdraw error message in the BillsPC state.", func(t *testing.T) {
		withdrawError := "withdraw error"
		billsPC := NewBillsPC(WithWithdrawError(withdrawError))
		if billsPC.withdrawError != withdrawError {
			t.Errorf("Expected the withdraw error message to be set in the BillsPC state, but it was not set.")
		}
	})
}
