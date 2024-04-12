package stonewall

// Detected time complexity:
// O(N)
// Score 100% https://app.codility.com/demo/results/trainingTFZQPC-E3D/
func Solution(H []int) int {
	stack := make([]int, 0)

	blockCounter := 0

	for i := range H {
		for len(stack) != 0 && stack[len(stack)-1] > H[i] {
			stack = stack[:len(stack)-1]
		}
		//if if len(stack) != 0 && stack[len(stack)-1] == H[i]  do nothing because you are in the same block
		if len(stack) == 0 || stack[len(stack)-1] != H[i] {
			stack = append(stack, H[i])
			blockCounter++
		}
	}
	return blockCounter
}
