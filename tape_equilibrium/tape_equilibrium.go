package tape_equilibrium

import (
	"math"
)

// Detected time complexity:
// O(N)
func Solution(A []int) int {
	if !IsValidArray(A) || !IsValidSize(A) {
		return 0
	}
	total := Sum(A)
	min := GetMin(A, total)
	return min
}

func IsValidSize(A []int) bool {
	if len(A) < 2 || len(A) > 100000 {
		return false
	}
	return true
}

func IsValidArray(A []int) bool {
	for i := range A {
		if A[i] < -1000 || A[i] > 1000 {
			return false
		}
	}
	return true
}

func GetMin(A []int, total int) int {
	leftSum := 0
	N := len(A)
	min := 1000 * N
	for i := 1; i < N; i++ {
		leftSum += A[i-1]
		rightSum := total - leftSum
		diff := int(math.Abs(float64(leftSum - rightSum)))

		if min > diff {
			min = diff
		}
	}
	return min
}

func Sum(A []int) (total int) {
	for i := range A {
		total += A[i]
	}
	return total
}
