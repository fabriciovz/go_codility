package dominator

import (
	"fmt"
	"sort"
)

// Detected time complexity:
// O(N*log(N)) or O(N)
func Solution(A []int) int {
	N := len(A)
	fmt.Println(A)
	if !IsvalidN(N) || !IsValidArray(A) {
		return -1
	}
	B := make([]int, len(A))
	copy(B, A)
	sort.Ints(B)
	Half := N / 2
	candidate := B[Half]
	fmt.Println(N)
	fmt.Println(candidate)
	count := 0
	lastIndexLedader := 0
	for i := 0; i < N; i++ {
		if A[i] == candidate {
			fmt.Printf("A[%d]=%d\n", i, A[i])
			count++
			lastIndexLedader = i
		}
	}
	if count > Half {
		return lastIndexLedader
	}
	return -1
}

func IsValidArray(A []int) bool {
	for i := range A {
		if !(A[i] >= -2147483648 && A[i] <= 2147483648) {
			return false
		}
	}
	return true
}

func IsvalidN(N int) bool {
	if N < 1 || N > 100000 {
		return false
	}
	return true
}
