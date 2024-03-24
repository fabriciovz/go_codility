package solution1

//Detected time complexity:
//O(N * max(Abs(A))**2)
//Minimum subset sum difference problem

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

	subset := make([]int, sum/2+1)
	subset[0] = 1
	for x, v := range countNum {
		if v == 0 {
			continue
		}
		for i := 0; i < len(subset); i++ {
			if subset[i] > 0 {
				subset[i] = v + 1
			} else if i >= x && subset[i-x] > 1 {
				subset[i] = subset[i-x] - 1
			}
		}
	}

	result := sum
	for k, s := range subset {
		if s == 0 {
			continue
		}
		if result > sum-k*2 {
			result = sum - k*2
		}
	}
	return result
}
