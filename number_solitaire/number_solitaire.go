package number_solitaire

const (
	MIN_VALUE = -10000
	MAX_VALUE = 10000
	MIN_INT   = -9223372036854775808
)

// Detected time complexity:
// O(N)
func Solution(A []int) int {
	N := len(A)
	if !IsvalidN(N) || !IsValidArray(A) {
		return N
	}
	R := make([]int, N)
	for i := 0; i < N; i++ {
		if i == 0 {
			R[0] = A[0]
		} else {
			currentMax := MIN_INT
			j := 0
			for j <= 6 {
				//for j := 1; j <= 6; j++ {
				if i-j >= 0 {
					currentMax = Max(currentMax, R[i-j]+A[i])
				}
			}
			R[i] = currentMax
		}
	}
	return R[N-1]
}

func Max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func IsValidArray(A []int) bool {
	for i := range A {
		if !(A[i] >= MIN_VALUE && A[i] <= MAX_VALUE) {
			return false
		}
	}
	return true
}

func IsvalidN(N int) bool {
	if N < 2 || N > 100000 {
		return false
	}
	return true
}
