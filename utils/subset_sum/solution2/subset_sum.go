package solution2

import "fmt"

// Dynamic programming
func Solution(A []int, N, Target int) bool {
	// Base case: If the target is 0, we've found a valid subset.
	if Target == 0 {
		return true
	}
	// Base case: If there are no elements left and the target is not 0, no subset exists.
	if N == 0 {
		return false
	}
	//Create an slice with cols = target + 1 and rows = N + 1
	cols := Target + 1
	rows := N + 1
	var dp = make([][]bool, rows)
	for i := range dp {
		dp[i] = make([]bool, cols)
	}

	// If sum is 0, then answer is true
	for i := 0; i <= N; i++ {
		dp[i][0] = true
	}

	// If sum is not 0 and set is empty,
	// then answer is false
	for i := 1; i <= Target; i++ {
		dp[0][i] = false
	}

	// Fill the subset table in bottom up manner
	for row := 1; row <= N; row++ {
		for col := 1; col <= Target; col++ {
			value := A[row-1]
			rowBack := row - 1
			if col < value {
				dp[row][col] = dp[rowBack][col]
			} else if col >= value {
				dp[row][col] = dp[rowBack][col] || dp[rowBack][col-value]
			}
		}
	}
	fmt.Println(dp)
	//Check if the target can be achieved by either including or excluding the last element.
	return dp[N][Target]
}
