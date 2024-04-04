package solution1

// Detected time complexity:
// O(N*log(N)) or O(N)
// 100% https://app.codility.com/demo/results/trainingAV665J-ATZ/
func Solution(A []int) int {
	N := len(A)

	if !IsvalidN(N) || !IsValidArray(A) {
		return -1
	}
	starts := make([]int, N)
	for i := 0; i < N; i++ {
		radius := A[i]
		startPos := i - radius

		if startPos < 0 {
			startPos = 0
		}
		starts[startPos]++
	}
	for i := 1; i < N; i++ {
		starts[i] += starts[i-1]
	}

	totalIntersections := 0
	for i := 0; i < N; i++ {
		radius := A[i]
		endPos := i + radius
		if endPos > N-1 {
			endPos = N - 1
		}
		intersections := Max(starts[i], starts[endPos]) - (i + 1)
		totalIntersections += intersections
		if totalIntersections > 10000000 {
			return -1
		}
	}

	return totalIntersections
}

func Max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func IsValidArray(A []int) bool {
	for i := range A {
		if !(A[i] >= 0 && A[i] <= 2147483647) {
			return false
		}
	}
	return true
}

func IsvalidN(N int) bool {
	if N < 0 || N > 100000 {
		return false
	}
	return true
}
