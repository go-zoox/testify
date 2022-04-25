package testify

import "testing"

func TestTypeOf(t *testing.T) {
	Type(t, "int", 1)
	Type(t, "string", "1")
}
