package testify

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ttacon/chalk"
)

// Equal logs error if expected and actual are not deeply equal.
func Equal[T comparable](t *testing.T, expected, received T, message ...string) {
	messageX := "expected and received are not equal"
	if len(message) > 0 {
		messageX = message[0]
	}

	if !reflect.DeepEqual(expected, received) {
		expectedX := fmt.Sprintf("%#v", expected)
		receivedX := fmt.Sprintf("%#v", received)
		t.Errorf(
			"%s\nexpected: %s\nreceived: %s",
			chalk.Bold.TextStyle(chalk.Red.Color(messageX)),
			chalk.Green.Color(expectedX),
			chalk.Red.Color(receivedX),
		)
	}
}
