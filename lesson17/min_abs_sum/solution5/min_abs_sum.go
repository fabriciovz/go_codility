package solution5

// Detected time complexity:
// O(N**2 * max(abs(A)))
// 72% https://app.codility.com/demo/results/trainingW8CDMR-QVE/
// Reference: https://codility.com/media/train/solution-min-abs-sum.pdf

func Abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func Solution(A []int) int {
	N := len(A)
	sum := 0
	if N == 0 {
		return sum
	}

	for i := 0; i < N; i++ {
		A[i] = Abs(A[i])
		sum += A[i]
	}

	cols := sum + 1
	var dp = make([]int, cols)

	dp[0] = 1
	for i := 0; i < N; i++ { //A elements
		for j := sum; j >= 0; j-- { //Current Target
			if dp[j] == 1 && j+A[i] <= sum {
				dp[j+A[i]] = 1
			}
		}
	}

	minDiff := sum
	middle := sum / 2
	for col := 1; col <= middle; col++ {
		if dp[col] == 1 {
			currentDiff := sum - 2*col
			if minDiff > currentDiff {
				minDiff = currentDiff
			}
		}
	}
	return minDiff
}
