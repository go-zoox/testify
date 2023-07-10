package testify

import (
	"testing"

	"github.com/ttacon/chalk"
)

// Assert logs error if not ok.
func Assert(t *testing.T, ok bool, message ...string) {
	messageX := "assert not ok"
	if len(message) > 0 {
		messageX = message[0]
	}

	if !ok {
		t.Errorf(
			"%s",
			chalk.Bold.TextStyle(chalk.Red.Color(messageX)),
		)
	}
}

// AssertTrue asserts that the specified value is true.
func AssertTrue(t *testing.T, received bool, message ...string) {
	messageX := "assert true but got false"
	if len(message) > 0 {
		messageX = message[0]
	}

	Assert(t, received, messageX)
}

// AssertFalse asserts that the specified value is false.
func AssertFalse(t *testing.T, received bool, message ...string) {
	messageX := "assert false but got true"
	if len(message) > 0 {
		messageX = message[0]
	}

	Assert(t, !received, messageX)
}
