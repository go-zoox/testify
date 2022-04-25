package testify

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ttacon/chalk"
)

// NotEqual logs error if expected and actual are equal.
func NotEqual[T comparable](t *testing.T, expected, received T, message ...string) {
	messageX := "expected and received are equal"
	if len(message) > 0 {
		messageX = message[0]
	}

	if reflect.DeepEqual(expected, received) {
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
