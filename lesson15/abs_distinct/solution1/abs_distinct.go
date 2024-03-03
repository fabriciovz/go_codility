package solution1

// Detected time complexity:
// O(N) or O(N*log(N))
func Solution(A []int) int {
	summaryA := map[int]int{}
	for i := range A {
		summaryA[Abs(A[i])]++
	}
	return len(summaryA)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
