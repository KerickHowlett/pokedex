package repl

import (
	"fmt"
	"testing"

	"github.com/KerickHowlett/pokedexcli/internal/tests/utils"
)

func TestWithPrintPrompt(t *testing.T) {
	fmt.Println("WithPrintPrompt should set the PrintPrompt field of the REPL struct")
	repl := setupREPLOptionTest(WithPrintPrompt(emptyFunctionMock))

	utils.ExpectSameEntity(t, repl.PrintPrompt, emptyFunctionMock, "PrintPrompt")
}
