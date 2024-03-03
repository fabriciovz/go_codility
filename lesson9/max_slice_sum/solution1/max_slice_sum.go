package solution1

import "fmt"

// Detected time complexity:
// O(N) or O(N**3)
// Brute force algorithm
func Solution(A []int) int {
	if !IsValidSize(A) || !IsValidArray(A) {
		fmt.Println("no valido")
		return 0
	}
	max := -1000000
	for i := 0; i < len(A); i++ {
		for j := i; j < len(A); j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += A[k]
			}
			if sum > max {
				max = sum
			}
		}
	}
	return max
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
