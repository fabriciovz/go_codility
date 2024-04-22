package solution2

import "fmt"

// Detected time complexity:
// O(N) or O(N*log(N))
func Solution(A []int) int {
	if len(A) == 1 {
		return A[0]
	}
	odd := CheckForOdd(A)
	return odd
}

func CheckForOdd(A []int) int {
	itemsMap := map[int]int{}

	// counting frequencies of array elements
	for i := 0; i < len(A); i++ {
		itemsMap[A[i]]++
	}
	fmt.Println(itemsMap)
	for i := range itemsMap {
		if itemsMap[i]%2 != 0 {
			return i
		}
	}
	return -1
}
