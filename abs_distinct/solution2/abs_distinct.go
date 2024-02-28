package solution2

// Caterpillar method
// Detected time complexity:
// O(N) or O(N*log(N))
func Solution(A []int) int {
	N := len(A)
	if !IsvalidN(N) {
		return 0
	}
	if N == 1 {
		return 1
	}
	myMax := Max(Abs(A[0]), Abs(A[N-1]))
	index_head := 0
	index_tail := N - 1
	c := 1
	for index_head <= index_tail {
		head := Abs(A[index_head])
		if head == myMax {
			index_head++
			continue
		}
		tail := Abs(A[index_tail])
		if tail == myMax {
			index_tail--
			continue
		}
		if head >= tail {
			myMax = head
			index_head++
		} else {
			myMax = tail
			index_tail--
		}
		c++
	}
	return c
}

func IsvalidN(N int) bool {
	if N < 1 || N > 100000 {
		return false
	}
	return true
}

func Max(x, y int) int {
	if x <= y {
		return y
	}
	return x
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
