package solution2

// Detected time complexity:
// O(log(N + M))
func Solution(N int, M int) int {
	return N / GCD(N, M, 1)
}

func GCD(A, B, res int) int {
	if A == B {
		return res * A
	} else if A%2 == 0 && B%2 == 0 {
		return GCD(A/2, B/2, 2*res)
	} else if A%2 == 0 {
		return GCD(A/2, B, res)
	} else if B%2 == 0 {
		return GCD(A, B/2, res)
	} else if A > B {
		return GCD(A-B, B, res)
	} else {
		return GCD(A, B-A, res)
	}
}
