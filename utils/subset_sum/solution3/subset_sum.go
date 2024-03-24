package solution3

// Dynamic programming
// for subset sum problem with space optimization to linear
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

	// Fill the subset table in bottom up manner

	for i := 1; i <= N; i++ { //A elements
		for j := 1; j <= Target; j++ { //Current Target
			if A[i-1] > j {
				curr[j] = prev[j]
			} else {
				curr[j] = prev[j] || prev[j-A[i-1]]
			}
		}
		// now curr becomes prev for row+1 th element
		copy(prev, curr)
	}
	//Check if the target can be achieved by either including or excluding the last element.
	return prev[Target]
}
