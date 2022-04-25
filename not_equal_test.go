package testify

import "testing"

func TestNotEqual(t *testing.T) {
	var a int = 1
	var b int = 2
	NotEqual(t, a, b)
	NotEqual(t, a, b, "a and b should be not equal")
}
