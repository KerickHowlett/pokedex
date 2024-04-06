package test_utils

import (
	"bytes"
	"io"
	"os"
)

// PrintStorage represents a storage for captured stdout prints.
//
// Fields:
//   - read: The IO reader used for reading the stdout.
//   - write: The IO writer used for writing to the stdout.
//   - old: A field to stash the original os.Stdout before mocking it's call signature.
type PrintStorage struct {
	read, write, old *os.File
}

// CaptureWithError captures the stdout printout of a function that returns an error.
// It returns the captured printout and the error returned by the function.
//
// Parameters:
//   - performTest: The function to perform the test.
//
// Returns:
//   - The captured printout.
//   - The error returned by the function.
//
// Example:
//
//	 printStorage := NewPrintStorage()
//	 printout, err := printStorage.CaptureWithError(func() error {
//		 fmt.Println("Hello, World!")
//		 return nil
//	 })
func (c *PrintStorage) CaptureWithError(performTest func() error) (printout string, err error) {
	c.captureStdOut()

	err = performTest()

	c.restoreStdOut()
	printout = c.getCapturedPrintOut()

	return printout, err
}

// Capture captures the stdout printout of a function.
// It returns the captured printout.
//
// Parameters:
//   - performTest: The function to perform the test.
//
// Returns:
//   - The captured printout.
func (c *PrintStorage) Capture(performTest func()) string {
	c.captureStdOut()

	performTest()

	c.restoreStdOut()

	return c.getCapturedPrintOut()
}

// captureStdOut redirects the standard output to a pipe, allowing it to be captured.
//
// Example:
//
//	printStorage := NewPrintStorage()
//	printStorage.captureStdOut()
func (c *PrintStorage) captureStdOut() {
	c.old = os.Stdout
	c.read, c.write, _ = os.Pipe()
	os.Stdout = c.write
}

// getCapturedPrintOut returns the captured printout.
//
// Returns:
//   - The captured printout.
//
// Example:
//
//	printStorage := NewPrintStorage()
//	printStorage.captureStdOut()
func (c *PrintStorage) getCapturedPrintOut() string {
	var buf bytes.Buffer
	io.Copy(&buf, c.read)

	return buf.String()
}

// restoreStdOut restores the standard output to the original os.Stdout.
//
// Example:
//
//	printStorage := NewPrintStorage()
//	printStorage.restoreStdOut()
func (c *PrintStorage) restoreStdOut() {
	c.write.Close()
	os.Stdout = c.old
}

// NewPrintStorage creates a new PrintStorage instance.
//
// Returns:
//   - A new instance of PrintStorage.
//
// Example:
//
//	printStorage := NewPrintStorage()
func NewPrintStorage() *PrintStorage {
	return &PrintStorage{}
}
