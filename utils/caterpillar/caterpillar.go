package binary_search

func Solution(A []int, S int) bool {
	N := len(A)
	front, total := 0, 0

	for back := 0; back < N; back++ {
		for front < N && total+A[front] <= S {
			total += A[front]
			front++
		}
		if total == S {
			return true
		}
		total -= A[back]
	}

	return false
}
