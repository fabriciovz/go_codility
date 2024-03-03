package solution2

import (
	"fmt"
)

// Detected time complexity:
// O(N * log(N)) or O(N ** 2)
// 88% https://app.codility.com/demo/results/trainingJK3M6B-QNR/
func Solution(A []int) []int {
	N := len(A)
	if !IsvalidN(N) {
		return []int{}
	}

	items := make(map[int]int, N)
	totalCount := make([]int, N)

	for i := 0; i < N; i++ {
		items[A[i]]++
	}
	for i := 0; i < N; i++ {
		divisors := 0
		nonDivisors := 0
		opp := 0
		for j := 1; j*j <= A[i]; j++ {
			if A[i]%j == 0 {
				val, ok := items[j]
				if ok {
					divisors += val
				}
				opp = A[i] / j
				if j*j < A[i] {
					val, ok = items[opp]
					if ok {
						divisors += items[opp]
					}
				}
			}
		}
		nonDivisors = N - divisors
		totalCount[i] = nonDivisors
	}
	fmt.Println(items)
	fmt.Println(totalCount)
	return totalCount
}

func IsvalidN(N int) bool {
	if N < 1 || N > 50000 {
		return false
	}
	return true
}
