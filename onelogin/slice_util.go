package onelogin

import (
	"slices"
	"strconv"
)

func sliceEqual(a, b []int) bool {
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
func sliceDiff(a, b []int) ([]int, []int) {
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

// intSliceToString converts a slice of ints to a string
func intSliceToString(s []int, sep string) string {
	if len(s) == 0 {
		return ""
	}

	var str string
	for _, v := range s {
		str += strconv.Itoa(v)
		str += sep
	}
	return str[:len(str)-len(sep)]
}
