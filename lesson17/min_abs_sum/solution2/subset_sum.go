package solution2

import "fmt"

// https://app.codility.com/demo/results/trainingBBRFCM-DN9/
// 72%
//Detected time complexity:
//O(N**2 * max(abs(A)))
//Ref: https://www.youtube.com/watch?v=FB0KUhsxXGY&t=110s

func Abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

// Dynamic programming
func Solution(A []int) int {
	N := len(A)
	// Base case: If there are no elements left and the target is not 0, no subset exists.
	if N == 0 {
		return 0
	}
	Target := 0
	for i := range A {
		A[i] = Abs(A[i])
		Target += A[i]
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
	minDiff := Target
	middle := Target / 2
	for col := 1; col <= middle; col++ {
		if dp[N][col] {
			currentDiff := Target - 2*col
			if minDiff > currentDiff {
				minDiff = currentDiff
			}
		}
	}

	PrintArray(dp)
	//Check if the target can be achieved by either including or excluding the last element.
	return minDiff
}

func PrintArray(dp [][]bool) {
	for row := 0; row < len(dp); row++ {
		fmt.Println(dp[row])
	}
}
