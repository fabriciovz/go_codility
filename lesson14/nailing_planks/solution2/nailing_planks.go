package solution2

// https://app.codility.com/demo/results/trainingYY5FCP-BX2/
func Solution(A []int, B []int, C []int) int {
	if !IsValidInput(A, B, C) {
		return -1
	}
	min := 1
	max := len(C)
	nail := make([]int, len(C)*2+1)
	total := -1
	for min <= max {
		miss := false
		mid := (min + max) / 2
		for i := 0; i < len(nail); i++ {
			nail[i] = 0
		}
		//filling nails up to mid
		for i := 0; i < mid; i++ {
			nail[C[i]]++
		}
		//doing prefix sum
		for i := 1; i < len(nail); i++ {
			nail[i] = nail[i] + nail[i-1]
		}
		for i := 0; i < len(A); i++ {
			if nail[A[i]-1] == nail[B[i]] {
				miss = true
			}
		}
		if miss {
			min = mid + 1
		} else {
			max = mid - 1
			total = mid
		}
	}
	return total
}

func IsValidInput(A []int, B []int, C []int) bool {
	lenA := len(A)
	lenB := len(B)
	lenC := len(C)

	if !(lenA == lenB) {
		return false
	}
	if lenA < 1 || lenA > 30000 {
		return false
	}
	if lenC < 1 || lenC > 30000 {
		return false
	}
	for i := 0; i < lenA; i++ {
		if !(A[i] <= B[i]) {
			return false
		}
		if A[i] < 1 || A[i] > 2*lenC {
			return false
		}
		if B[i] < 1 || B[i] > 2*lenC {
			return false
		}
	}
	for i := 0; i < lenC; i++ {
		if C[i] < 1 || C[i] > 2*lenC {
			return false
		}
	}
	return true
}
