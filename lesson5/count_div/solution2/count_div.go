package solution2

// Detected time complexity:
// O(1)
func Solution(A int, B int, K int) int {
	N := B - A
	if N < 0 || !IsvalidK(K) {
		return 0
	}
	// Add 1 explicitly as A is divisible by M
	if A%K == 0 {
		return B/K - A/K + 1
	}
	// A is not divisible by M
	return B/K - A/K
}

func IsvalidK(N int) bool {
	if N < 1 || N > 2000000000 {
		return false
	}
	return true
}
