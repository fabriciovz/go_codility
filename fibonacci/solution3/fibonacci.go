package solution3

// Fibonacci using dynamic programming
// The time complexity O(n)
// Tabulation
func Solution(N int) int {
	if N <= 1 {
		return N
	}
	fibSerie := make([]int, N+1)
	fibSerie[0] = 0
	fibSerie[1] = 1
	for i := 2; i <= N; i++ {
		fibSerie[i] = fibSerie[i-2] + fibSerie[i-1]
	}
	return fibSerie[N]
}
