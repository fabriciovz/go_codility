package min_max_division

//Detected time complexity:
//O(N*log(N+M))

func Solution(K int, M int, A []int) int {
	N := len(A)

	if !IsValidNK(K) || !IsValidNK(N) || !IsValidM(M) || !IsValidArray(A) {
		return -1
	}
	minSum, maxSum := 0, 0
	for i := range A {
		maxSum += A[i]
		minSum = Max(minSum, A[i])
	}
	if K == N {
		return minSum
	}
	result := maxSum
	for minSum <= maxSum {
		mid := (minSum + maxSum) / 2
		if IsSplitValid(mid, K, A) {
			//find smaller mid
			maxSum = mid - 1
			result = mid
		} else {
			//find larger mid
			minSum = mid + 1
		}

	}
	return result
}

func IsSplitValid(mid, K int, A []int) bool {
	sum := 0
	splits := 0
	for i := range A {
		if sum+A[i] <= mid {
			sum += A[i]
		} else {
			sum = A[i]
			splits++
		}
	}
	return splits < K
}

func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func IsValidM(N int) bool {
	if N < 0 || N > 10000 {
		return false
	}
	return true
}

func IsValidNK(N int) bool {
	if N < 1 || N > 100000 {
		return false
	}
	return true
}

func IsValidArray(A []int) bool {
	for i := range A {
		if A[i] < 0 || A[i] > 10000 {
			return false
		}
	}
	return true
}
