package repl

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	utils "testtools/utils"
)

// @SECTION: TEST CASES

func TestDefaultNewREPL(t *testing.T) {
	fmt.Println("NewREPL should create a new REPL instance with the provided options")

	repl := NewREPL(WithCommandExecutor(commandExecutorMock))

	if repl.Scanner == bufio.NewScanner(os.Stdin) {
		t.Errorf("repl.Scanner is not the default bufio.NewScanner(os.Stdin) instance")
	}

	utils.ExpectSameEntity(t, repl.CommandExecutor, commandExecutorMock, "CommandExecutor")
	utils.ExpectSameEntity(t, repl.CommandExecutor, commandExecutorMock, "CommandExecutor")
	utils.ExpectSameEntity(t, repl.PrintNewLine, defaultPrintNewLine, "PrintNewLine")
	utils.ExpectSameEntity(t, repl.PrintPrompt, defaultPrintPrompt, "PrintPrompt")
}

func TestNewREPL_PanicIfCommandExecutorNotSet(t *testing.T) {
	fmt.Println("NewREPL should panic if CommandExecutor is not set")
	defer func() {
		if recover := recover(); recover == nil {
			t.Errorf("Expected NewREPL to panic, but it didn't")
		}
	}()

	NewREPL()
}

// @SECTION: MOCKS

func commandExecutorMock(args string) error {
	return nil
}

func emptyFunctionMock() {}

func setupREPLOptionTest(option REPLOption) *REPL {
	return NewREPL(WithCommandExecutor(commandExecutorMock), option)
}
