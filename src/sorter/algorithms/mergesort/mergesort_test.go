//mergesort_test.go
package mergesort

import (
	"testing"
)

func TestMerge(t *testing.T) {
	left := []int{1, 2, 3}
	right := []int{4, 5, 6}
	merged := merge(left, right)
	if merged[0] != 1 || merged[1] != 2 || merged[2] != 3 || merged[3] != 4 || merged[4] != 5 || merged[5] != 6 {
		t.Error("Merge Failed.", merged, "Expected 1 2 3 4 5 6")
	}
}

func TestMergeSort(t *testing.T) {
	array := []int{3, 2, 5, 7, 9, 1, 4}
	sorted := MergeSort(array)
	if sorted[0] != 1 || sorted[1] != 2 || sorted[2] != 3 || sorted[3] != 4 || sorted[4] != 5 || sorted[5] != 7 || sorted[6] != 9 {
		t.Error("Merge Sort Failed.", sorted, "Expected 1 2 3 4 5 7 9")
	}
}
