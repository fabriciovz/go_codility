package solution2

// Detected time complexity:
// O(N + M)
func Solution(N int, A []int) []int {
	if !IsvalidN(N) || !IsvalidN(len(A)) || !IsValidArray(A, N) {
		return []int{}
	}
	counters := make([]int, N)
	maxCounter := 0
	startLine := 0

	for i := range A {
		if A[i] <= N {
			counters[A[i]-1] = Max(counters[A[i]-1]+1, startLine+1)
			if maxCounter < counters[A[i]-1] {
				maxCounter = counters[A[i]-1]
			}
		} else if A[i] == N+1 {
			startLine = maxCounter
		}
	}
	for i := 0; i < N; i++ {
		counters[i] = Max(counters[i], startLine)
	}
	return counters
}

func Max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func IsValidArray(A []int, N int) bool {
	for i := range A {
		if A[i] < 1 || A[i] > N+1 {
			return false
		}
	}
	return true
}

func IsvalidN(N int) bool {
	if N < 1 || N > 100000 {
		return false
	}
	return true
}
