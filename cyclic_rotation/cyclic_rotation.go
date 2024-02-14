package cyclic_rotation

func Solution(A []int, K int) []int {
	if IsvalidK(K) && IsValidArray(A) {
		return CyclicRotation(A, K)
	}
	return A
}

func CyclicRotation(A []int, K int) []int {
	if K == 1 {
		return Rotate(A)
	}
	A = Rotate(A)
	return CyclicRotation(A, K-1)
}

func Rotate(A []int) []int {
	var newArray []int
	newArray = make([]int, len(A))
	arrayLen := len(A)
	for i := range A {
		if i == 0 {
			newArray[i] = A[arrayLen-i-1]
		} else {
			newArray[i] = A[i-1]
		}
	}
	return newArray
}

func IsvalidK(K int) bool {
	if !(K >= 0 && K <= 100) || K == 0 {
		return false
	}
	return true
}

func IsValidArray(A []int) bool {
	for i := range A {
		if !(A[i] >= -1000 && A[i] <= 1000) {
			return false
		}
	}
	return true
}
