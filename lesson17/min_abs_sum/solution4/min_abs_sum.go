package solution3

// Detected time complexity:
// O(N**2 * max(abs(A)))
// 72% https://app.codility.com/demo/results/training6EFCQ6-ZD8/
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

	//Get the Max item in array A and convert the array to abs
	A[0] = Abs(A[0])
	for i := 0; i < N; i++ {
		A[i] = Abs(A[i])
		sum += A[i]
	}

	cols := sum + 1
	var prev = make([]bool, cols)

	// If sum is 0, then answer is true
	prev[0] = true
	// If sum is not 0 and set is empty,
	// then answer is false
	for i := 1; i < cols; i++ {
		prev[i] = false
	}

	// curr array to store the current row result generated
	// with help of prev array
	var curr = make([]bool, cols)
	curr[0] = true

	for i := 1; i <= N; i++ { //A elements
		for j := 1; j <= sum; j++ { //Current Target
			if A[i-1] > j {
				curr[j] = prev[j]
			} else {
				curr[j] = prev[j] || prev[j-A[i-1]]
			}
		}
		copy(prev, curr)
	}

	minDiff := sum
	middle := sum / 2
	for col := 1; col <= middle; col++ {
		if prev[col] {
			currentDiff := sum - 2*col
			if minDiff > currentDiff {
				minDiff = currentDiff
			}
		}
	}
	return minDiff
}
