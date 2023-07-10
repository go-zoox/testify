package testify

import (
	"fmt"
	"testing"

	"github.com/ttacon/chalk"
)

// Enums logs error if element is not one of enums.
func Enums[T comparable](t *testing.T, enums []T, element T, message ...string) {
	messageX := "value is not in enum"
	if len(message) > 0 {
		messageX = message[0]
	}

	if !arrayIncludes(enums, element) {
		expectedX := fmt.Sprintf("%#v", enums)
		receivedX := fmt.Sprintf("%#v", element)
		t.Errorf(
			"%s\nenums: %s\nvalue: %s",
			chalk.Bold.TextStyle(chalk.Red.Color(messageX)),
			chalk.Green.Color(expectedX),
			chalk.Red.Color(receivedX),
		)
	}
}
