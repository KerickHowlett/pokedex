package bills_pc

import (
	"testing"

	p "pokemon"
	f "test_tools/fixtures"
	"test_tools/utils"
)

func TestBillsPC_Deposit(t *testing.T) {
	runDepositTests := func() (billsPC *BillsPC, pokemon *p.Pokemon) {
		pokemon = &p.Pokemon{Name: f.PokemonName}
		billsPC = &BillsPC{caughtPokemon: &CaughtPokemon{}}
		billsPC.Deposit(pokemon)
		return billsPC, pokemon
	}

	t.Run("should add a Pokemon to the BillsPC if it has not already been caught.", func(t *testing.T) {
		t.Parallel()
		billsPC, pokemon := runDepositTests()
		if _, found := (*billsPC.caughtPokemon)[pokemon.Name]; !found {
			t.Errorf("Expected to find the caught Pokemon in the BillsPC state, but it was not found.")
		}
	})

	t.Run("should not add a Pokemon to the BillsPC if it has already been caught.", func(t *testing.T) {
		t.Parallel()
		billsPC, pokemon := runDepositTests()

		// Attempt to deposit the same Pokemon again.
		duplicatePokemon := &p.Pokemon{Name: pokemon.Name}
		billsPC.Deposit(duplicatePokemon)
		actualPokemon := (*billsPC.caughtPokemon)[pokemon.Name]

		// Ensure that the duplicate Pokemon was not added to the BillsPC.
		utils.ExpectSameEntity(t, pokemon, actualPokemon, "Pokemon")
	})
}

func TestBillsPC_Inspect(t *testing.T) {
	runInspectTests := func() (billsPC *BillsPC, expectedPokemon *p.Pokemon) {
		expectedPokemon = &p.Pokemon{Name: f.PokemonName}
		billsPC = &BillsPC{caughtPokemon: &CaughtPokemon{expectedPokemon.Name: expectedPokemon}}
		return billsPC, expectedPokemon
	}

	t.Run("should return the Pokemon with the provided name if it has been caught.", func(t *testing.T) {
		t.Parallel()
		billsPC, expectedPokemon := runInspectTests()
		if actualPokemon, _ := billsPC.Inspect(expectedPokemon.Name); actualPokemon != expectedPokemon {
			t.Errorf("Expected to find the caught Pokemon in the BillsPC state, but it was not found.")
		}
	})

	t.Run("should return a boolean indicating whether the Pokemon has been caught.", func(t *testing.T) {
		t.Parallel()
		billsPC, expectedPokemon := runInspectTests()
		if _, found := billsPC.Inspect(expectedPokemon.Name); !found {
			t.Errorf("Expected to find the caught Pokemon in the BillsPC state, but it was not found.")
		}
	})
}

func TestBillsPC_TotalCaughtPokemon(t *testing.T) {
	runTotalCaughtPokemonTests := func() (billsPC *BillsPC, expectedTotalCaughtPokemon int) {
		expectedTotalCaughtPokemon = 1
		billsPC = &BillsPC{caughtPokemon: &CaughtPokemon{
			f.PokemonName: &p.Pokemon{Name: f.PokemonName},
		}}
		return billsPC, expectedTotalCaughtPokemon
	}

	t.Run("should return the total number of Pokemon caught by the user.", func(t *testing.T) {
		t.Parallel()
		billsPC, expectedTotalCaughtPokemon := runTotalCaughtPokemonTests()
		if actualTotalCaughtPokemon := billsPC.TotalCaughtPokemon(); actualTotalCaughtPokemon != expectedTotalCaughtPokemon {
			t.Errorf("Expected to find the caught Pokemon in the BillsPC state, but it was not found.")
		}
	})
}

func TestBillsPC_Withdraw(t *testing.T) {
	runWithdrawTests := func(pokemonName string) (actualPokemon *p.Pokemon, expectedPokemon *p.Pokemon, expectedErrorMessage string, err error) {
		expectedPokemon = &p.Pokemon{Name: f.PokemonName}
		billsPC := &BillsPC{caughtPokemon: &CaughtPokemon{expectedPokemon.Name: expectedPokemon}}
		actualPokemon, err = billsPC.Withdraw(pokemonName)
		return actualPokemon, expectedPokemon, billsPC.withdrawError, err
	}

	t.Run("Given the user has caught the Pokemon possessing the provided name", func(t *testing.T) {
		t.Parallel()
		t.Run("Then it should remove the target Pokemon from the BillsPC.", func(t *testing.T) {
			if actualPokemon, expectedPokemon, _, _ := runWithdrawTests(f.PokemonName); actualPokemon != expectedPokemon {
				t.Errorf("Expected to find the caught Pokemon in the BillsPC state, but it was not found.")
			}
		})
	})

	t.Run("Given the user has NOT caught the Pokemon the possessing the provided name", func(t *testing.T) {
		t.Parallel()
		t.Run("should return an error if the Pokemon was not found in the BillsPC.", func(t *testing.T) {
			t.Parallel()
			if _, _, _, err := runWithdrawTests(f.Invalid); err == nil {
				t.Errorf("Expected to receive an error when withdrawing a Pokemon that was not found in the BillsPC, but no error was returned.")
			}
		})

		t.Run("should return an error message when a Pokemon is not found in the BillsPC.", func(t *testing.T) {
			t.Parallel()
			if _, _, expectedErrorMessage, err := runWithdrawTests(f.Invalid); err.Error() != expectedErrorMessage {
				t.Errorf("Expected to receive the error message when withdrawing a Pokemon that was not found in the BillsPC, but a different error message was returned.")
			}
		})
	})
}
