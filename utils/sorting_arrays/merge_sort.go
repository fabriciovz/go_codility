package sorting_arrays

func MergeSort(A []int) []int {
	sorted := Sort(A)
	return sorted
}

func Sort(A []int) []int {
	if len(A) < 2 {
		return A
	}
	left, right := Split(A)
	left = Sort(left)
	right = Sort(right)
	merge := Merge(left, right)
	return merge
}

// get left and right sorted slices
func Merge(left []int, right []int) []int {
	merge := make([]int, len(left)+len(right))
	l, r := 0, 0
	for i := 0; i < len(merge); i++ {
		if l >= len(left) {
			merge[i] = right[r]
			r++
			continue
		} else if r >= len(right) {
			merge[i] = left[l]
			l++
			continue
		}
		if left[l] <= right[r] {
			merge[i] = left[l]
			l++
		} else {
			merge[i] = right[r]
			r++
		}
	}
	return merge
}

func Split(A []int) ([]int, []int) {
	left, right := A[0:(len(A)/2)], A[(len(A)/2):]
	return left, right
}
