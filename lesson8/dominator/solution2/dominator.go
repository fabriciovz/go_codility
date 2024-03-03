package solution2

// Detected time complexity:
// O(N*log(N)) or O(N)
func Solution(A []int) int {
	N := len(A)
	if !IsvalidN(N) || !IsValidArray(A) {
		return -1
	}
	Half := N / 2
	summary := map[int][]int{}
	for i := 0; i < N; i++ {
		summary[A[i]] = append(summary[A[i]], i)
		if len(summary[A[i]]) > Half {
			return i
		}
	}
	return -1
}

func IsValidArray(A []int) bool {
	for i := range A {
		if !(A[i] >= -2147483648 && A[i] <= 2147483648) {
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
