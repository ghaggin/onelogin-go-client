package util

import "slices"

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	slices.Sort(a)
	slices.Sort(b)

	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

// Diff returns a not in b and b not in a
//
// Note: this could be optimized
func Diff(a, b []int) ([]int, []int) {
	var aNotB, bNotA []int

	for _, v := range a {
		if !slices.Contains(b, v) {
			aNotB = append(aNotB, v)
		}
	}

	for _, v := range b {
		if !slices.Contains(a, v) {
			bNotA = append(bNotA, v)
		}
	}

	return aNotB, bNotA
}
