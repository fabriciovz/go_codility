package solution1

// Detected time complexity:
// O(N)
func Solution(N int) int {
	divisorCounter := 0
	for i := 1; i <= N; i++ {
		if N%i == 0 {
			divisorCounter++
		}
	}
	return divisorCounter
}
