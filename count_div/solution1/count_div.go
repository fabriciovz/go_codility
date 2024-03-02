package solution1

// Detected time complexity:
// O(B-A)
// https://app.codility.com/demo/results/training6UYXJ7-W6H/
func Solution(A int, B int, K int) int {
	N := B - A
	if N < 0 || !IsvalidK(K) {
		return 0
	}
	prefixSum := make([]int, N+1)

	if A%K == 0 {
		prefixSum[0] = 1
	}

	for i, j := A+1, 1; i <= B; i, j = i+1, j+1 {
		if i%K == 0 {
			prefixSum[j] = prefixSum[j-1] + 1
		} else {
			prefixSum[j] += prefixSum[j-1]
		}
	}
	return prefixSum[N]
}

func IsvalidK(N int) bool {
	if N < 1 || N > 2000000000 {
		return false
	}
	return true
}
