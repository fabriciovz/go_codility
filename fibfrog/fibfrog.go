package solution2

import "fmt"

func Solution(A []int) int {
	max_distance := len(A) + 1
	fibonacci := make([]int, max_distance)
	//jumps := make([]int, max_distance)

	fibonacci[0] = 0
	fibonacci[1] = 1
	for i := 2; i < max_distance; i++ {
		fibonacci[i] = fibonacci[i-2] + fibonacci[i-1]
	}

	fmt.Println(fibonacci)
	return 0
}
