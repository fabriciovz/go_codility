package max_non_overlapping_segments

// Detected time complexity:
// O(N)
func Solution(A []int, B []int) int {
	N := len(A)
	if N <= 1 {
		return N
	}
	if !IsValidArray(A, B) || !IsvalidN(N) {
		return 0
	}
	currentPosition := -1
	segmentCount := 0
	for i := 0; i < N; i++ {
		if A[i] > currentPosition {
			currentPosition = B[i]
			segmentCount++
		}
	}
	return segmentCount
}

func IsvalidN(N int) bool {
	if N < 0 || N > 30000 {
		return false
	}
	return true
}

func IsValidArray(A []int, B []int) bool {
	for i := range A {
		if A[i] > B[i] {
			return false
		}
	}
	return true
}
