package testify

import (
	"testing"
)

// True asserts that the specified value is true.
func True(t *testing.T, received bool, message ...string) {
	messageX := "expected true but got false"
	if len(message) > 0 {
		messageX = message[0]
	}

	Equal(t, true, received, messageX)
}
