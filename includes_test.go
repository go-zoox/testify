package testify

import "testing"

func TestIncludes(t *testing.T) {
	var list []int = []int{1, 2, 3}
	var element int = 3
	Includes(t, list, element)
	Includes(t, list, element, "list is not include element")
}
