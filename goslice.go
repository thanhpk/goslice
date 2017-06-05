package slice

// Superset return true id a is superset of b
func Superset(a []string, b []string) bool {
	for _, i := range b {
		found := false
		for _, j := range a {
			if i == j {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Equals compare two slice, do not care about order
func Equals(a []string, b []string) bool {
	return Superset(a, b) && Superset(b, a)
}
