package testify

import "testing"

func TestRange(t *testing.T) {
	Range(t, 0, 10, 1)
	// Range(t, 0, 10, 11)
}
