package solution3

//Detected time complexity:
//O(N * max(abs(A))**2)
//100% https://app.codility.com/demo/results/training4XQ5MZ-5W4/

func Abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
func Min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func Solution(A []int) int {
	N := len(A)
	sum := 0

	if N == 0 {
		return sum
	}

	//Get the Max item in array A and convert the array to abs
	maxA := Abs(A[0])
	A[0] = Abs(A[0])
	for i := 1; i < N; i++ {
		A[i] = Abs(A[i])
		if A[i] > maxA {
			maxA = A[i]
		}
	}

	countNum := make([]int, maxA+1)
	for _, a := range A {
		countNum[a] += 1
		sum += a
	}

	dp := make([]int, sum+1)
	//fill dp array
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	dp[0] = 0

	for i := range countNum {
		if countNum[i] == 0 {
			continue
		}
		for j := 0; j < len(dp); j++ {
			if dp[j] >= 0 {
				dp[j] = countNum[i]
			} else if j >= i && dp[j-i] > 0 {
				dp[j] = dp[j-i] - 1
			}
		}
	}

	result := sum
	rang := sum/2 + 1
	for i := 1; i < rang; i++ {
		if dp[i] >= 0 {
			result = Min(result, sum-2*i)
		}
	}
	return result
}
