package distinct

// O(N*log(N)) or O(N)
func Solution(A []int) int {
	if len(A) > 100000 || !IsValidArrayElem(A) {
		return 0
	}
	sumary := map[int]int{}
	for i := range A {
		sumary[A[i]]++
	}
	return len(sumary)
}

func IsValidArrayElem(A []int) bool {
	for i := 0; i < len(A); i++ {
		if A[i] < -1000000 || A[i] > 1000000 {
			return false
		}
	}
	return true
}
