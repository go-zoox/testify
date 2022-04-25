package testify

import (
	"fmt"
	"testing"

	"github.com/go-zoox/core-utils/array"
	"github.com/ttacon/chalk"
)

// Range logs error if value is not in range [min, max].
func Range[T array.Number](t *testing.T, min T, max T, value T, message ...string) {
	messageX := fmt.Sprintf("value is not in range (min: %v, max: %v)", min, max)
	if len(message) > 0 {
		messageX = message[0]
	}

	if value > max || value < min {
		expectedX := fmt.Sprintf("%#v", []T{min, max})
		receivedX := fmt.Sprintf("%#v", value)
		t.Errorf(
			"%s\nexpected: %s\nreceived: %s",
			chalk.Bold.TextStyle(chalk.Red.Color(messageX)),
			chalk.Green.Color(expectedX),
			chalk.Red.Color(receivedX),
		)
	}
}
