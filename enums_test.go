package testify

import "testing"

func TestEnums(t *testing.T) {
	var enums []int = []int{1, 2, 3}
	var element int = 3
	Enums(t, enums, element)
	Enums(t, enums, element, "element is not one of enums")
}
