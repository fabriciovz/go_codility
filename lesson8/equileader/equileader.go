package equileader

// Detected time complexity:
// O(N)
func Solution(A []int) int {
	N := len(A)
	Half := N / 2
	summary := map[int][]int{}
	dominator := 0
	for i := 0; i < N; i++ {
		summary[A[i]] = append(summary[A[i]], i)
		if len(summary[A[i]]) > Half {
			dominator = A[i]
		}
	}
	leftCounter := 0
	rightCounter := N
	equileaders := 0
	leadersInRightSide := len(summary[dominator])
	leadersInLeftSide := 0
	for i := 0; i < N; i++ {
		if A[i] == dominator {
			leadersInLeftSide++
			leadersInRightSide--
		}
		leftCounter++
		rightCounter--
		if leadersInLeftSide > leftCounter/2 {
			if leadersInRightSide > rightCounter/2 {
				equileaders++
			}
		}
	}
	return equileaders
}
