package frog_river_one

// Detected time complexity:
// O(N)
func Solution(X int, A []int) int {
	if !IsValidArrayElem(X, A) {
		return -1
	}
	if !IsValidRange(len(A)) || !IsValidRange(X) {
		return -1
	}
	return Check(X, A)
}

func IsValidRange(size int) bool {
	if size < 1 || size > 100000 {
		return false
	}
	return true
}

func IsValidArrayElem(X int, A []int) bool {
	for i := 0; i < len(A); i++ {
		if A[i] < 1 || A[i] > X {
			return false
		}
	}
	return true
}

func Check(X int, A []int) int {
	steps := map[int]int{}
	for i := range A {
		_, ok := steps[A[i]]
		if !ok && A[i] <= X {
			steps[A[i]] = i
		}
	}
	max := 0
	for i := 1; i <= X; i++ {
		_, ok := steps[i]
		if !ok {
			return -1
		}
		if max < steps[i] {
			max = steps[i]
		}
	}
	return max
}
