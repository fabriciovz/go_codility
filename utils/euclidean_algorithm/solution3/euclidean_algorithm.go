package solution3

// Binary Euclidean algorithm
// Greatest common divisor
func Solution(A, B, res int) int {
	if A == B {
		return res * A
	} else if A%2 == 0 && B%2 == 0 {
		return Solution(A/2, B/2, 2*res)
	} else if A%2 == 0 {
		return Solution(A/2, B, res)
	} else if B%2 == 0 {
		return Solution(A, B/2, res)
	} else if A > B {
		return Solution(A-B, B, res)
	} else {
		return Solution(A, B-A, res)
	}
}
