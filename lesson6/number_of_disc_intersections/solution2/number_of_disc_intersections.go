package solution2

import "sort"

// Detected time complexity:
// O(N**2)
// 81% https://app.codility.com/demo/results/trainingGQCBCV-A8B/
func Solution(A []int) int {
	N := len(A)

	if !IsvalidN(N) || !IsValidArray(A) {
		return -1
	}

	starts := make([]int, N)
	ends := make([]int, N)
	results := make([]int, N)

	for i := 0; i < N; i++ {
		start := i - A[i]
		end := i + A[i]
		starts[i] = start
		ends[i] = end
	}

	sort.Ints(starts)

	for i := 0; i < N; i++ {
		count := 0
		for j := 0; j < N; j++ {
			if ends[i] >= starts[j] {
				count++
			}
		}
		results[i] = count - 1
	}
	intersections := 0
	for i := 0; i < N; i++ {
		intersections += (results[i] - i)
		if intersections > 10000000 {
			return -1
		}
	}
	return intersections
}

func IsValidArray(A []int) bool {
	for i := range A {
		if !(A[i] >= 0 && A[i] <= 2147483647) {
			return false
		}
	}
	return true
}

func IsvalidN(N int) bool {
	if N < 0 || N > 100000 {
		return false
	}
	return true
}
