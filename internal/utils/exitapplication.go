package utils

import "os"

const (
	OK = 0
)

func ExitApplication(statusCode int) error {
	os.Exit(statusCode)

	return nil
}
