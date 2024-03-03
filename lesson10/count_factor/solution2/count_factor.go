package solution1

import (
	"fmt"
	"math"
)

// Detected time complexity:
// O(sqrt(N))
func Solution(N int) int {
	sqrt := math.Sqrt(float64(N))
	sqrtInt := int(sqrt)
	fmt.Printf("Square root of %d is :%d\n", N, sqrtInt)
	factorCounter := 0
	for i := 1; i <= sqrtInt; i++ {
		fmt.Printf("i is :%d\n", i)
		if N%i == 0 {
			factorCounter += 2
		}
	}
	if N == sqrtInt*sqrtInt {
		factorCounter--
	}
	return factorCounter
}
