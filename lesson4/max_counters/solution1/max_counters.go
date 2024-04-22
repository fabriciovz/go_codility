package solution1

// 77% https://app.codility.com/demo/results/trainingFXF4JD-JGR/
func Solution(N int, A []int) []int {
	if !IsvalidN(N) || !IsvalidN(len(A)) || !IsValidArray(A, N) {
		return []int{}
	}
	counters := make([]int, N)
	maxCounter := 0
	for i := range A {
		if A[i] >= 1 && A[i] <= N {
			counters[A[i]-1]++
			if maxCounter < counters[A[i]-1] {
				maxCounter = counters[A[i]-1]
			}
		} else if A[i] == N+1 {
			for j := 0; j < N; j++ {
				counters[j] = maxCounter
			}

		}
	}
	return counters
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
