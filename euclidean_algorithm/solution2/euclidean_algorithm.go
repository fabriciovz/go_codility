package solution2

// Euclidean algorithm by division
func Solution(A, B int) int {
	if B == 0 {
		return A
	} else {
		return Solution(B, A%B)
	}
}
