package solution2

import "fmt"

// Detected time complexity:
// O(N)
func Solution(A []int) int {
	if !IsValidSize(A) || !IsValidArray(A) {
		fmt.Println("no valido")
		return 0
	}
	currentSum := 0
	maxSum := -1000000
	for i := 0; i < len(A); i++ {
		currentSum += A[i]
		if currentSum < A[i] {
			currentSum = A[i]
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

func IsValidSize(A []int) bool {
	if len(A) < 1 || len(A) > 1000000 {
		return false
	}
	return true
}

func IsValidArray(A []int) bool {
	for i := range A {
		if A[i] < -1000000 || A[i] > 1000000 {
			return false
		}
	}
	return true
}
