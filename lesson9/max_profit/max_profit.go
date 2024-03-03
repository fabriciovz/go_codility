package max_profit

// Detected time complexity:
// O(N)
func Solution(A []int) int {
	maxProfit := 0
	cumulativeProfit := 0
	for i := 1; i < len(A); i++ {
		dayProfit := A[i] - A[i-1]
		cumulativeProfit += dayProfit
		if cumulativeProfit > 0 {
			if cumulativeProfit > maxProfit {
				maxProfit = cumulativeProfit
			}
		} else {
			cumulativeProfit = 0
		}
	}
	return maxProfit
}
