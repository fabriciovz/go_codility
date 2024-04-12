package stonewall

// Detected time complexity:
// O(N)
// Score 100 https://app.codility.com/demo/results/training2DF9U3-KYA/
func Solution(H []int) int {
	stack := make([]int, 0)

	blockCounter := 0

	for i := range H {
		for len(stack) != 0 && stack[len(stack)-1] > H[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 && stack[len(stack)-1] == H[i] {

		} else {
			stack = append(stack, H[i])
			blockCounter++
		}
	}
	return blockCounter
}
