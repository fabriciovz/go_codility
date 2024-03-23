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

func Max(A []int) int {
	maxN := Abs(A[0])
	for i := 1; i < len(A); i++ {
		absN := Abs(A[i])
		if absN > maxN {
			maxN = absN
		}
	}
	return maxN
}

func Solution(A []int) int {
	maxLen := 0

	if len(A) == 0 {
		return maxLen
	}

	var maxA = Max(A)
	countNum := make([]int, maxA+1)
	for _, a := range A {
		absNum := Abs(a)
		countNum[absNum] = countNum[absNum] + 1
		maxLen = maxLen + absNum
	}

	subset := make([]int, maxLen/2+1)
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

	result := maxLen
	for k, s := range subset {
		if s == 0 {
			continue
		}
		if result > maxLen-k*2 {
			result = maxLen - k*2
		}
	}
	return result
}
