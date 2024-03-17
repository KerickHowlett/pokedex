package repl

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStartREPL(t *testing.T) {
	fmt.Println("StartREPL should execute command and prompt for input until user exits")

	scannerMock := &ScannerMock{isEnabled: mockedScannerEnabledStatus, userInput: "exit"}

	executedCommand := ""
	mockedCommandExecutor := func(args string) error {
		executedCommand = args
		defer func() { scannerMock.isEnabled = false }()
		return nil
	}

	repl := NewREPL(
		WithCommandExecutor(mockedCommandExecutor),
		WithPrintPrompt(emptyFunctionMock),
		WithScanner(scannerMock),
		WithParseInput(parserMock),
		WithPrintNewLine(emptyFunctionMock),
	)

	repl.StartREPL()

	expectedCommand := scannerMock.userInput
	if !reflect.DeepEqual(executedCommand, expectedCommand) {
		t.Errorf("Expected executed commands to be %v, but got %v", expectedCommand, executedCommand)
	}
}
