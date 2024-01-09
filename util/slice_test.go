package util

import "testing"

func TestDiff(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}

	aNotB, bNotA := Diff(a, b)

	if !Equal(aNotB, []int{1}) {
		t.Errorf("aNotB should be [1], got %v", aNotB)
	}

	if !Equal(bNotA, []int{4}) {
		t.Errorf("bNotA should be [4], got %v", bNotA)
	}

	a = []int{1, 2, 3}
	b = []int{1, 2, 3}
	aNotB, bNotA = Diff(a, b)
	if len(aNotB) != 0 {
		t.Errorf("aNotB should be empty, got %v", aNotB)
	}
	if len(bNotA) != 0 {
		t.Errorf("bNotA should be empty, got %v", bNotA)
	}
}
