package binary_search

func Solution(A []int, X int) int {
	N := len(A)
	low := 0
	high := N - 1
	for low <= high {
		//mid := (low + high) / 2
		mid := low + (high-low)/2
		if A[mid] < X {
			low = mid + 1
		} else if A[mid] > X {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
