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
