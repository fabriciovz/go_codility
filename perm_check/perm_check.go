package perm_check

import "sort"

// Detected time complexity:
// O(N) or O(N * log(N))
func Solution(A []int) int {
	N := len(A)
	if !IsvalidN(N) || !IsValidArrayElem(A) {
		return 0
	}
	sort.Ints(A)
	if A[0] != 1 {
		return 0
	}
	checkCounter := 0
	for i := 0; i < N-1; i++ {
		if A[i] == A[i+1]-1 {
			checkCounter++
		}
	}
	if checkCounter == N-1 {
		return 1
	} else {
		return 0
	}
}

func IsvalidN(N int) bool {
	if N < 1 || N > 100000 {
		return false
	}
	return true
}

func IsValidArrayElem(A []int) bool {
	for i := 0; i < len(A); i++ {
		if A[i] < 1 || A[i] > 1000000 {
			return false
		}
	}
	return true
}
