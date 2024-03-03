package solution1

func Solution(N int) int {
	if N <= 1 {
		return N
	} else {
		return Solution(N-2) + Solution(N-1)
	}
}
