package array

// EqualElements tells whether a and b contain the same elements. Elements does not need to be in same index
// A nil argument is equivalent to an empty slice.
func EqualElements(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for _, v := range a {
		containsV := false
		for _, w := range b {
			if w == v {
				containsV = true
				break
			}
		}
		if !containsV {
			return false
		}
	}
	return true
}
