package tie_ropes

// Detected time complexity:
// O(N)
// https://app.codility.com/demo/results/trainingRCKCRF-MB9/
func Solution(K int, A []int) int {
	acum := 0
	result := 0
	for i := 0; i < len(A); i++ {
		acum += A[i]
		if acum/K >= 1 {
			result++
			acum = 0
		}
	}
	return result
}
