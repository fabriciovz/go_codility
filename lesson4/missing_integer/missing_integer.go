package missing_integer

import (
	"sort"
)

// Detected time complexity:
// O(N) or O(N * log(N))
// 100% https://app.codility.com/demo/results/trainingCPVPDC-SDT/
func Solution(A []int) int {
	N := len(A)
	sort.Ints(A)

	if A[N-1] < 1 || N == 0 {
		return 1
	}
	countA := make([]int, A[N-1]+1)

	for i := 0; i < N; i++ {
		if A[i] > 0 {
			countA[A[i]]++
		}
	}
	minIndex := A[N-1] + 1
	for i := 1; i < len(countA); i++ {
		if countA[i] == 0 {
			minIndex = i
			break
		}
	}
	return minIndex
}
