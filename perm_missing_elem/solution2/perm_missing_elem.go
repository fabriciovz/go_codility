package solution2

import "fmt"

// Detected time complexity:
// O(N) or O(N * log(N))
func Solution(A []int) int {
	sizeA := len(A)
	if sizeA == 0 {
		return 1
	}
	acum := 0
	for i := range A {
		acum += A[i]
	}

	//I added 1 because there is a "missing element"
	n := sizeA + 1
	expectedAcum := (n * (n + 1)) / 2

	fmt.Println(acum)
	fmt.Println(expectedAcum)

	return expectedAcum - acum
}

//1,2,3,5 = 11
//1,2,3,4,5 = 15
// 15 - 11 = 4; 4 is the missing element
