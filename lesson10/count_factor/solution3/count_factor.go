package solution1

// Detected time complexity:
// O(sqrt(N))
func Solution(N int) int {
	factorCounter := 0
	i := 1
	for i*i < N {
		if N%i == 0 {
			factorCounter += 2
		}
		i++
	}
	if N == i*i {
		factorCounter++
	}
	return factorCounter
}
