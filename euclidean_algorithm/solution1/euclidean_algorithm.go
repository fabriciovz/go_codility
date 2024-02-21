package solution1

// original implementation of Euclidean's algorithm
// Euclidean algorithm by subtraction
func Solution(A, B int) int {
	if A == B {
		return A
	}
	if A > B {
		return Solution(A-B, B)
	} else {
		return Solution(A, B-A)
	}
}
