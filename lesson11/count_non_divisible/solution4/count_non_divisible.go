package solution4

// Detected time complexity:
// O(N * log(N))
// https://app.codility.com/demo/results/trainingZJUV7R-HBH/
func Solution(A []int) []int {
	N := len(A)
	if !IsvalidN(N) {
		return []int{}
	}
	totalCount := make([]int, N)
	divisors := make([]int, (N*2)+1)
	nonDivisors := map[int]int{}
	for i := 0; i < N; i++ {
		divisors[A[i]]++
		nonDivisors[A[i]] = 0
	}
	for v, _ := range nonDivisors {
		divCount := 0
		for j := 1; j*j <= v; j++ {
			if v%j == 0 {
				divCount += divisors[j]
				if v/j != j {
					divCount += divisors[v/j]
				}
			}
		}
		nonDivisors[v] = N - divCount
	}
	for i := 0; i < N; i++ {
		totalCount[i] = nonDivisors[A[i]]
	}
	return totalCount
}

func IsvalidN(N int) bool {
	if N < 1 || N > 50000 {
		return false
	}
	return true
}
