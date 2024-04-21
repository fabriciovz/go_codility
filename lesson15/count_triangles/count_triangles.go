package count_triangles

import "sort"

// Detected time complexity:
// O(N**2)
// 100% score https://app.codility.com/demo/results/trainingY2Q75U-48R/
// Solution based on https://www.youtube.com/watch?v=6Kg_hS6rG8k
// and https://www.geeksforgeeks.org/find-number-of-triangles-possible/
func Solution(A []int) int {
	N := len(A)
	sort.Ints(A)
	triangles := 0
	for back := 0; back < N-2; back++ {
		// Initialize index of the rightmost third
		// element
		front := back + 2
		// Fix the second element
		for m := back + 1; m < N-1; m++ {
			for front < N && A[back]+A[m] > A[front] {
				front++
			}
			// Total number of possible triangles that can
			// be formed with the two fixed elements is
			// front - m - 1. The two fixed elements are arr[back]
			// and arr[m]. All elements between arr[m+1]/ to
			// arr[front-1] can form a triangle with arr[back] and
			// arr[m]. One is subtracted from front because front is
			// incremented one extra in above while loop. front
			// will always be greater than m. If m becomes
			// equal to front, then above loop will increment front,
			// because arr[front] + arr[back] is always greater than arr[front]
			if front > m {
				triangles += front - m - 1
			}
		}
	}
	return triangles
}
