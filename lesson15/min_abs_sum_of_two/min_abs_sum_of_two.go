package min_abs_sum_of_two

import "sort"

// Detected time complexity:
// O(N * log(N))
// 100% score https://app.codility.com/demo/results/trainingJUCBZE-2PH/
func Solution(A []int) int {
	sort.Ints(A)
	N := len(A)

	back := 0
	front := N - 1
	minAbsSum := Abs(A[back] + A[front])

	for back <= front {
		currentSum := A[back] + A[front]
		currentAbsSum := Abs(A[back] + A[front])
		minAbsSum = Min(minAbsSum, currentAbsSum)
		// If the sum has become
		// positive, we should know that the head can be moved left
		if currentSum <= 0 {
			back++
		} else {
			front--
		}
	}
	return minAbsSum
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
