package solution6

import "math"

//Detected time complexity:
//O(N * max(abs(A))**2)
//100% https://app.codility.com/demo/results/trainingE6HS25-26E/
// Reference: https://www.youtube.com/watch?v=0IL_LohRtiQ&t=570s

func Abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func Solution(A []int) int {
	N := len(A)
	sum := 0
	if N == 0 {
		return sum
	}

	mapA := map[int]int{}
	for i := 0; i < N; i++ {
		A[i] = Abs(A[i])
		sum += A[i]
		mapA[A[i]]++
	}

	var middle = float64(sum) / 2.0
	var rang = int(math.Floor(middle))
	cols := rang + 1
	var dp = make([]int, cols)
	dp[0] = 0
	for i := 1; i < cols; i++ {
		dp[i] = -1
	}

	for val := range mapA { //A elements
		for j := 0; j < cols; j++ { //Current Target
			if dp[j] >= 0 {
				dp[j] = mapA[val]
			} else if j-val >= 0 {
				dp[j] = dp[j-val] - 1
			}
		}
	}

	for i := rang; i >= 0; i-- {
		if dp[i] >= 0 {
			return int((middle - float64(i)) * 2.0)
		}
	}
	return -1
}
