package solution1

import "fmt"

//Detected time complexity:
//O(n^2)

func Solution(A []int) int {
	if len(A) == 1 {
		return A[0]
	}
	sorted := Sort(A)
	fmt.Println(sorted)
	odd := CheckForOdd(sorted)
	return odd
}

func CheckForOdd(A []int) int {
	odd := 0
	for i := 0; i < len(A); i++ {
		count := 0
		for j := 0; j < len(A); j++ {
			if A[i] == A[j] {
				count++
			}
		}
		if count%2 != 0 {
			return A[i]
		}
	}
	return odd
}

func Sort(A []int) []int {
	aux := 0
	for j := 0; j < len(A); j++ {
		for i := 0; i < len(A)-1; i++ {
			if A[i] > A[i+1] {
				aux = A[i]
				A[i] = A[i+1]
				A[i+1] = aux
			}
		}
	}
	return A
}
