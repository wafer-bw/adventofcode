package set

// AllUnique returns true if all elements in the slice are unique.
func AllUnique(list []string) bool {
	m := map[string]struct{}{}
	for _, v := range list {
		if _, ok := m[v]; ok {
			return false
		}
		m[v] = struct{}{}
	}
	return true
}
