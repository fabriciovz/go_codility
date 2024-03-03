package sorting_arrays

// Quadratic time O(n^2)
// small data set: OK
func SelectionSort(A []int) []int {
	for i := 0; i < len(A)-1; i++ {
		min := i
		for j := i + 1; j < len(A); j++ {
			if A[min] > A[j] {
				min = j
			}
		}
		temp := A[i]
		A[i] = A[min]
		A[min] = temp
	}
	return A
}
