package solution3

import (
	"fmt"
)

// Detected time complexity:
// O(N * log(N)) or O(N ** 2)
// 88% https://app.codility.com/demo/results/trainingMJB5V7-3AV/
func Solution(A []int) []int {
	N := len(A)
	if !IsvalidN(N) {
		return []int{}
	}
	totalCount := make([]int, N)
	divisors := make([]int, (N*2)+1)
	for i := 0; i < N; i++ {
		divisors[A[i]]++
	}
	for i := 0; i < N; i++ {
		divCount := 0
		for j := 1; j*j <= A[i]; j++ {
			if A[i]%j == 0 {
				divCount += divisors[j]
				if A[i]/j != j {
					divCount += divisors[A[i]/j]
				}
			}
		}
		totalCount[i] = N - divCount
	}
	fmt.Println(divisors)
	fmt.Println(totalCount)
	return totalCount
}

func IsvalidN(N int) bool {
	if N < 1 || N > 50000 {
		return false
	}
	return true
}
