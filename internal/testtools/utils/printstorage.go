package utils

import (
	"bytes"
	"io"
	"os"
)

type PrintStorage struct {
	read, write, old *os.File
}

func (c *PrintStorage) CaptureWithError(performTest func() error) (printout string, err error) {
	c.captureStdOut()

	err = performTest()

	c.restoreStdOut()
	printout = c.getCapturedPrintOut()

	return printout, err
}

func (c *PrintStorage) Capture(performTest func()) string {
	c.captureStdOut()

	performTest()

	c.restoreStdOut()

	return c.getCapturedPrintOut()
}

func (c *PrintStorage) captureStdOut() {
	c.old = os.Stdout
	c.read, c.write, _ = os.Pipe()
	os.Stdout = c.write
}

func (c *PrintStorage) getCapturedPrintOut() string {
	var buf bytes.Buffer
	io.Copy(&buf, c.read)

	return buf.String()
}

func (c *PrintStorage) restoreStdOut() {
	c.write.Close()
	os.Stdout = c.old
}

func NewPrintStorage() *PrintStorage {
	return &PrintStorage{}
}
