package solution1

// Backtracking/recursion
func Solution(A []int, N, Target int) bool {
	// Base case: If the target is 0, we've found a valid subset.
	if Target == 0 {
		return true
	}
	// Base case: If there are no elements left and the target is not 0, no subset exists.
	if N == 0 {
		return false
	}

	//If the last element is greater than the target, exclude it.
	if A[N-1] > Target {
		return Solution(A, N-1, Target)
	}
	//Check if the target can be achieved by either including or excluding the last element.
	return Solution(A, N-1, Target) || Solution(A, N-1, Target-A[N-1])
}
