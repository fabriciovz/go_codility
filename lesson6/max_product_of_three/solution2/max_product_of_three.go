package solution1

import (
	"fmt"
	"sort"
)

// Time Complexity: O(n log n), due to sorting.
// Space Complexity: O(log n), Sorting stack space using merge sort or quicksort.
func Solution(A []int) int {
	sort.Ints(A)
	fmt.Println(A)

	N := len(A)
	maxProd1 := A[N-1] * A[N-2] * A[N-3]
	maxProd2 := A[N-1] * A[0] * A[1]

	max := 0
	if maxProd1 > maxProd2 {
		max = maxProd1
	} else {
		max = maxProd2
	}
	return max
}
