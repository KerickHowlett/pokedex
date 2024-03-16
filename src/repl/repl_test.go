package repl

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"testing"
)

// @SECTION: TEST CASES

func TestDefaultNewREPL(t *testing.T) {
	fmt.Println("NewREPL should create a new REPL instance with the provided options")
	repl := NewREPL(
		WithCommandExecutor(commandExecutorMock),
	)

	expectSameEntity(t, repl.CommandExecutor, commandExecutorMock, "CommandExecutor")
	expectSameEntity(t, repl.PrintNewLine, printNewLine, "PrintNewLine")
	expectSameEntity(t, repl.PrintPrompt, printPrompt, "PrintPrompt")
	expectSameEntity(t, repl.Scanner, bufio.NewScanner(os.Stdin), "Scanner")
}

func TestNewREPL_PanicIfCommandExecutorNotSet(t *testing.T) {
	fmt.Println("NewREPL should panic if CommandExecutor is not set")
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected NewREPL to panic, but it didn't")
		}
	}()

	NewREPL()
}

func TestWithCommandExecutor(t *testing.T) {
	fmt.Println("WithFields should set the Fields field of the Sanitizer struct")
	repl := NewREPL(WithCommandExecutor(commandExecutorMock))

	expectSameEntity(t, repl.CommandExecutor, commandExecutorMock, "CommandExecutor")
}

func TestWithParseInput(t *testing.T) {
	fmt.Println("WithParseInput should set the ParseInput field of the REPL struct")
	repl := setupREPLOptionTest(WithParseInput(parserMock))

	expectSameEntity(t, repl.ParseInput, parserMock, "ParseInput")
}

func TestWithPrintNewLine(t *testing.T) {
	fmt.Println("WithPrintNewLine should set the PrintNewLine field of the REPL struct")
	repl := setupREPLOptionTest(WithPrintNewLine(emptyFunctionMock))

	expectSameEntity(t, repl.PrintNewLine, emptyFunctionMock, "PrintNewLine")
}

func TestWithPrintPrompt(t *testing.T) {
	fmt.Println("WithPrintPrompt should set the PrintPrompt field of the REPL struct")
	repl := setupREPLOptionTest(WithPrintPrompt(emptyFunctionMock))

	expectSameEntity(t, repl.PrintPrompt, emptyFunctionMock, "PrintPrompt")
}

func TestWithScanner(t *testing.T) {
	fmt.Println("WithScanner should set the Scanner field of the REPL struct")
	scannerMock := &ScannerMock{}
	repl := setupREPLOptionTest(WithScanner(scannerMock))

	expectSameEntity(t, repl.Scanner, scannerMock, "Scanner")
}

// @SECTION: ASSERTERS

func expectSameEntity(t *testing.T, actual any, expected any, fieldName string) {
	if reflect.ValueOf(actual).Pointer() != reflect.ValueOf(expected).Pointer() {
		t.Errorf("Expected %s to be set with the argued function, but instead got %v\n", fieldName, actual)
	}
}

// @SECTION: MOCKS

func commandExecutorMock(args string) error {
	return nil
}

func emptyFunctionMock() {}

type ScannerMock struct {
	userInput string
	isEnabled bool
}

var mockedScannerEnabledStatus bool = true

func (s *ScannerMock) Scan() bool {
	return s.isEnabled
}

func (s *ScannerMock) Text() string {
	return s.userInput
}

func parserMock(input string) []string {
	return []string{input}
}

// @SECTION: FACTORIES

func setupREPLOptionTest(option REPLOption) *REPL {
	return NewREPL(WithCommandExecutor(commandExecutorMock), option)
}
