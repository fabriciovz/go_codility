package fish

const (
	Downstrem = 1
)

func Solution(A []int, B []int) int {
	N1 := len(A)
	N2 := len(B)
	if !IsvalidN(N1) || N1 != N2 || !IsValidArrayA(A) || !IsValidArrayB(B) {
		return 0
	}
	var stack []int
	passFish := 0
	for i, direction := range B {
		fish := A[i]
		if direction == Downstrem {
			stack = append(stack, fish)
		} else {
			for len(stack) > 0 {
				curFish := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				//if A[i] > curFish {
				if curFish > fish {
					stack = append(stack, curFish)
					break
				}
			}

			if len(stack) == 0 {
				passFish = passFish + 1
			}
		}
	}
	return len(stack) + passFish
}

func IsvalidN(N int) bool {
	if N < 1 || N > 100000 {
		return false
	}
	return true
}

func IsValidArrayA(A []int) bool {
	checker := map[int]int{}
	for i := 0; i < len(A); i++ {
		if A[i] < 0 || A[i] > 1000000000 {
			return false
		}
		checker[A[i]]++
	}
	if len(checker) != len(A) {
		return false
	}
	return true
}

func IsValidArrayB(A []int) bool {
	for i := 0; i < len(A); i++ {
		if A[i] < 0 || A[i] > 1 {
			return false
		}
	}
	return true
}
