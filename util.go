package testify

// Includes returns true if the array contains the value.
func arrayIncludes[K comparable](values []K, value K) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}

	return false
}
