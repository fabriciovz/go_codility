package min_avg_two_slice

import "math"

// Detected time complexity:
// O(N)
// 100% https://app.codility.com/demo/results/trainingRB8HSK-WE5/
func Solution(A []int) int {
	prefixSum := make([]int, len(A)+1)
	for i := 1; i <= len(A); i++ {
		prefixSum[i] = A[i-1] + prefixSum[i-1]
	}
	minVal := float64(10000)
	minIndex := 0

	//Math background problem
	//slices of 4 or more numbers always contain a smaller slice with lower average
	//Reference:https://www.youtube.com/watch?v=Xu_hTjFAauk

	for i := 1; i < len(prefixSum)-2; i++ {
		val1 := float64(prefixSum[i+1]-prefixSum[i-1]) / 2.0
		val2 := float64(prefixSum[i+2]-prefixSum[i-1]) / 3.0

		if minVal > val1 || minVal > val2 {
			minVal = math.Min(val1, val2)
			minIndex = i - 1
		}
	}
	val1 := float64(prefixSum[len(prefixSum)-1]-prefixSum[len(prefixSum)-3]) / 2.0

	if minVal > val1 {
		return len(prefixSum) - 3
	}

	return minIndex
}
