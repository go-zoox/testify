package testify

import "testing"

func TestEqual(t *testing.T) {
	var a int = 1
	var b int = 1
	Equal(t, a, b)
	Equal(t, a, b, "a and b should be equal")
}
