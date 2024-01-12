package onelogin

import "testing"

func TestDiff(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}

	aNotB, bNotA := sliceDiff(a, b)

	if !sliceEqual(aNotB, []int{1}) {
		t.Errorf("aNotB should be [1], got %v", aNotB)
	}

	if !sliceEqual(bNotA, []int{4}) {
		t.Errorf("bNotA should be [4], got %v", bNotA)
	}

	a = []int{1, 2, 3}
	b = []int{1, 2, 3}
	aNotB, bNotA = sliceDiff(a, b)
	if len(aNotB) != 0 {
		t.Errorf("aNotB should be empty, got %v", aNotB)
	}
	if len(bNotA) != 0 {
		t.Errorf("bNotA should be empty, got %v", bNotA)
	}
}

func (s *OneLoginTestSuite) Test_intSliceToString() {
	s.Equal(intSliceToString([]int{1, 2, 3}, ","), "1,2,3")
	s.Equal(intSliceToString([]int{1, 2, 3}, ""), "123")
	s.Equal(intSliceToString([]int{1, 2, 3}, " "), "1 2 3")
	s.Equal(intSliceToString([]int{1, 2, 3}, "x"), "1x2x3")
	s.Equal(intSliceToString([]int{}, "x"), "")
	s.Equal(intSliceToString([]int{}, ""), "")
	s.Equal(intSliceToString([]int{}, ","), "")
}
