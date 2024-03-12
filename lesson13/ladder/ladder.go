package ladder

import (
	"math"
)

// Detected time complexity:
// O(L)
// https://app.codility.com/demo/results/trainingUYMNR9-7VW/
func Solution(A []int, B []int) []int {
	N := len(A)
	if !IsvalidL(N) || !IsValidArrayA(A, N) || !IsValidArrayB(B) {
		return []int{}
	}
	var result []int
	var fibonacci []int

	fibonacci = append(fibonacci, 0)
	fibonacci = append(fibonacci, 1)
	fibonacci = append(fibonacci, 2)
	//Create a fibo list
	for i := 3; i < N+1; i++ {
		pow := int(math.Pow(2, 30))
		//To avoid overflow, it is applied the pisano period to repeat fibo sequence
		//https://en.wikipedia.org/wiki/Pisano_period
		next := (fibonacci[i-2] + fibonacci[i-1]) % pow
		fibonacci = append(fibonacci, next)
	}
	for i := 0; i < N; i++ {
		pow := int(math.Pow(2, float64(B[i])))
		fibo := fibonacci[A[i]]
		element := fibo % pow
		result = append(result, element)
	}

	return result
}

func IsvalidL(N int) bool {
	if N < 1 || N > 50000 {
		return false
	}
	return true
}

func IsValidArrayA(A []int, L int) bool {
	for i := range A {
		if A[i] < 1 || A[i] > L {
			return false
		}
	}
	return true
}

func IsValidArrayB(B []int) bool {
	for i := range B {
		if B[i] < 1 || B[i] > 30 {
			return false
		}
	}
	return true
}
