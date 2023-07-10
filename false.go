package testify

import (
	"testing"
)

// False asserts that the specified value is false.
func False(t *testing.T, received bool, message ...string) {
	messageX := "expected false but got true"
	if len(message) > 0 {
		messageX = message[0]
	}

	Equal(t, false, received, messageX)
}
