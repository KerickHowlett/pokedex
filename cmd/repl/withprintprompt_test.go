package repl

import (
	"fmt"
	"reflect"
	"testing"

	u "testtools/utils"
)

func TestWithPrintPrompt(t *testing.T) {
	fmt.Println("WithPrintPrompt should set the PrintPrompt field of the REPL struct")

	repl := setupREPLOptionTest(WithPrintPrompt(emptyFunctionMock))

	if !reflect.DeepEqual(repl.PrintPrompt, emptyFunctionMock) {
		fmt.Printf("PrintPrompt: %p\n", repl.PrintPrompt)
		fmt.Printf("PrintPrompt: %p\n", emptyFunctionMock)
	}

	u.ExpectSameEntity(t, repl.PrintPrompt, emptyFunctionMock, "PrintPrompt")
}
