package solution2

// Detected time complexity:
// O(1)
func Solution(X int, Y int, D int) int {
	steps := 1
	if X >= Y || D < 1 {
		return 0
	}
	if (Y-X)%D == 0 {
		return (Y - X) / D
	} else {
		return (Y-X)/D + 1
	}

	return steps
}
