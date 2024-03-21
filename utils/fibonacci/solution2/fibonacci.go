package solution2

func Solution(N int, memo []int) int {
	if N <= 1 {
		return N
	} else if memo[N] == 0 {
		memo[N] = Solution(N-2, memo) + Solution(N-1, memo)
	}
	return memo[N]
}
