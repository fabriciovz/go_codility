package solution3

import "math"

// 100% https://app.codility.com/demo/results/trainingERFWXA-UUQ/
const (
	MAX_INT = 1000000000
)

func Solution(N int) int {
	if N < 1 || N > MAX_INT {
		return 0
	}

	for i := math.Sqrt(float64(N)); i >= 1; i-- {
		index := int(i)
		if N%index == 0 {
			return 2 * (index + N/index)
		}
	}
	return 0
}
