package solution1

import (
	"fmt"
	"sort"
)

// Detected time complexity:
// O(N * log(N)) or O(N ** 2)
// 88% score
// // https://app.codility.com/demo/results/trainingKW8NXC-AZR/
func Solution(A []int) []int {
	N := len(A)
	if !IsvalidN(N) {
		return []int{}
	}
	B := make([]int, N)
	totalCount := make([]int, N)
	copy(B, A)
	sort.Ints(B)
	items := make([]int, B[N-1]+1)
	for i := 0; i < N; i++ {
		items[A[i]]++
	}
	for i := 0; i < N; i++ {
		divisors := 0
		nonDivisors := 0
		opp := 0
		for j := 1; j*j <= A[i]; j++ {
			if A[i]%j == 0 {
				divisors += items[j]
				opp = A[i] / j
				if j*j < A[i] {
					divisors += items[opp]
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
