package solution1

import "fmt"

// Detected time complexity:
// O(N**3)
// Brute Force Approach
func Solution(A []int) int {
	if !IsValidArray(A) || !IsvalidN(len(A)) {
		return 0
	}
	if len(A) == 3 {
		return A[0] * A[1] * A[2]
	}
	max := -1000
	for i := 0; i < len(A); i++ {
		for j := 1; j < len(A); j++ {
			for k := 2; k < len(A); k++ {
				if j > i && k > j {
					prod := A[i] * A[j] * A[k]
					fmt.Printf("(%d,%d,%d) = %d\n", i, j, k, prod)
					if prod > max {
						max = prod
					}
				}
			}
		}
	}
	return max
}

func IsValidArray(A []int) bool {
	for i := range A {
		if !(A[i] >= -1000 && A[i] <= 1000) {
			return false
		}
	}
	return true
}

func IsvalidN(N int) bool {
	if N < 3 || N > 100000 {
		return false
	}
	return true
}
