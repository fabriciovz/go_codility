package solution1

import "fmt"

//Detected time complexity:
//O(N ** 2)

func Solution(A []int) int {
	if len(A) == 0 {
		return 1
	}
	if len(A) == 1 {
		if A[0] == 1 {
			return 2
		} else {
			return 1
		}
	}
	sorted := Sort(A)
	missing := CheckMissing(sorted)
	fmt.Println(sorted)
	fmt.Println(missing)

	return missing
}

func CheckMissing(A []int) int {
	count := 1
	for i := 0; i < len(A); i++ {
		fmt.Printf("A[%d]=%d\n", i, A[i])
		fmt.Printf("count=%d\n", count)
		if A[i] != count {
			return count
		}
		count++
	}
	return count
}

func Sort(A []int) []int {
	aux := 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A)-1; j++ {
			if A[j] > A[j+1] {
				aux = A[j]
				A[j] = A[j+1]
				A[j+1] = aux
			}
		}
	}
	return A
}
