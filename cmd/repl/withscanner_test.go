package repl

import (
	"fmt"
	"testing"

	m "github.com/KerickHowlett/pokedexcli/internal/tests/mocks"
	"github.com/KerickHowlett/pokedexcli/internal/tests/utils"
)

func TestWithScanner(t *testing.T) {
	fmt.Println("WithScanner should set the Scanner field of the REPL struct")

	scanner := m.NewScannerMock()
	repl := setupREPLOptionTest(WithScanner(scanner))

	utils.ExpectSameEntity(t, repl.Scanner, scanner, "Scanner")
}
