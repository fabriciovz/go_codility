package caterpillar_method

func Solution(A []int, S int) bool {
	N := len(A)
	front, total := 0, 0
	for back := range A {
		for front < N && total+A[front] <= S {
			total += A[front]
			front++
		}
		if total == S {
			return true
		} else {
			total -= A[back]
			return false
		}
	}
	return false
}
