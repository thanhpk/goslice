package slice

import (
	"strings"
)

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
func Equal(a []string, b []string) bool {
	return Superset(a, b) && Superset(b, a)
}

// ContainString check if s contains e
func ContainString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Merge merge a and b into a big slice. O(n^2)
func Merge(s1 []string, s2 []string) []string {
	ret := make([]string, len(s1))
	copy(ret, s1)
	for _, s := range s2 {
		found := false
		for _, subs := range s1 {
			if subs == s {
				found = true
				break
			}
		}
		if !found {
			ret = append(ret, s)
		}
	}
	return ret
}


// Substract return all elements which is in slice but not in subslice. O(n^2)
func Substract(slice []string, subslice []string) []string {
	ret := make([]string, 0)
	for _, s := range slice {
		found := false
		for _, subs := range subslice {
			if subs == s {
				found = true
				break
			}
		}
		if !found {
			ret = append(ret, s)
		}
	}
	return ret
}

func Find(slice []string, ele string) int {
	for i, e := range slice {
		if e == ele {
			return i
		}
	}
	return -1
}

func Trim(slice []string) []string {
	newslice := make([]string, 0)
	for _, s := range slice {
		strim := strings.Trim(s, " ")
		if strim != "" {
			newslice = append(newslice, strim)
		}
	}
	return newslice
}
