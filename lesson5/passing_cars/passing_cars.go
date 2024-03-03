package passing_cars

func Solution(A []int) int {
	nEast := 0
	cross := 0
	for i := 0; i < len(A); i++ {
		if A[i] == 0 {
			nEast++
		} else {
			cross += nEast
		}
		if cross > 1000000000 {
			return -1
		}
	}
	return cross
}
