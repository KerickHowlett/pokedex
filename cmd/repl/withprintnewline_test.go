package repl

import (
	"fmt"
	"testing"

	"github.com/KerickHowlett/pokedexcli/internal/tests/utils"
)

func TestWithPrintNewLine(t *testing.T) {
	fmt.Println("WithPrintNewLine should set the PrintNewLine field of the REPL struct")
	repl := setupREPLOptionTest(WithPrintNewLine(emptyFunctionMock))

	utils.ExpectSameEntity(t, repl.PrintNewLine, emptyFunctionMock, "PrintNewLine")
}
