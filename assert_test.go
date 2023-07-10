package testify

import "testing"

func TestAssert(t *testing.T) {
	var a int = 1
	var b int = 2
	Assert(t, a != b)
	Assert(t, a != b, "a and b should be not equal")
}

func TestAssertTrue(t *testing.T) {
	AssertTrue(t, true)
}

func TestAssertFalse(t *testing.T) {
	AssertFalse(t, false)
}
