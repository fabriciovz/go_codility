package max_double_slice_sum

// Detected time complexity:
// O(N)
// 100% socre https://app.codility.com/demo/results/trainingE3F55B-FNK/
// Solution based on Kadane's Algorithm
func Solution(A []int) int {
	if !IsValidSize(A) {
		return 0
	}
	N := len(A)
	lr := make([]int, N)
	rl := make([]int, N)
	s := 0
	//filling left right array
	for i := 1; i < N-1; i++ {
		s += A[i]
		if s < 0 {
			s = 0
		}
		lr[i] = s
	}
	//filling right left array
	s = 0
	for i := N - 2; i > 0; i-- {
		s += A[i]
		if s < 0 {
			s = 0
		}
		rl[i] = s
	}
	//filling right left array
	maxSum := 0
	for i := 0; i < N-2; i++ {
		maxSum = Max(maxSum, lr[i]+rl[i+2])
	}
	return maxSum
}

func IsValidSize(A []int) bool {
	if len(A) < 3 {
		return false
	}
	return true
}

func Max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

/*
func IsValidArray(A []int) bool {
	for i := range A {
		if A[i] < -1000000 || A[i] > 1000000 {
			return false
		}
	}
	return true
}*/
