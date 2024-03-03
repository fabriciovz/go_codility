package solution2

const (
	IntMax = 100000
)

// Detected time complexity:
// O(N * log(N))
// https://app.codility.com/demo/results/trainingUCKT7H-NRQ/
func Solution(A []int) int {
	N := len(A)
	if !IsvalidN(N) || !IsValidArray(A) {
		return -1
	}
	if N == 0 || N == 1 || N == 2 {
		return 1
	}
	var fibonacci []int
	jumps := make([]int, N+1)

	fibonacci = append(fibonacci, 1)
	fibonacci = append(fibonacci, 2)
	//Create a fibo list
	next := 3
	fibIndex := 3
	for next <= N+1 {
		fibonacci = append(fibonacci, next)
		next = fibonacci[fibIndex-2] + fibonacci[fibIndex-1]
		fibIndex++
	}
	//First fill all the jumps with IntMax
	for i := 0; i <= N; i++ {
		jumps[i] = IntMax // valid values will always be <=N<INT32_MAX bc 1 is a Fib number
	}
	//Initial jump
	for _, f := range fibonacci {
		if f <= N+1 {
			jumps[f-1] = 1
		}
	}
	// 1D dijkstra’s
	for start := 0; start < N; start++ {
		if A[start] == 1 {
			minJumps := jumps[start]
			if minJumps != IntMax {
				for _, j := range fibonacci {
					if start+j > N {
						break
					}
					//sum of fibo num + array index is equal to N (it's arrived to the end
					//the next element is a leaf??
					if (start+j == N) || (A[start+j] == 1) {
						//fmt.Printf("jumps[%d] = min(jumps[%d], %d)\n", start+j, start+j, minJumps+1)
						jumps[start+j] = min(jumps[start+j], minJumps+1)
					}
				}
			}
		}
	}
	if jumps[N] == IntMax {
		return -1
	} else {
		return jumps[N]
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func IsvalidN(N int) bool {
	if N < 0 || N > IntMax {
		return false
	}
	return true
}

func IsValidArray(A []int) bool {
	for i := range A {
		if A[i] != 0 && A[i] != 1 {
			return false
		}
	}
	return true
}
