package solution1

// Detected time complexity:
// O(Y-X)
func Solution(X int, Y int, D int) int {
	distance := X
	steps := 1
	if X >= Y || D < 1 {
		return 0
	}
	for {
		distance += D
		//fmt.Printf("distance=%d steps=%d\n", distance, steps)
		if distance >= Y {
			break
		}
		steps++
	}
	return steps
}
