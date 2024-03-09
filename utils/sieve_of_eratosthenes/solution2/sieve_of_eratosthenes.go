package solution2

import "math"

// Find all prime numbers till N
func Solution(N int) []int {
	prime := make([]bool, N+1)
	for i := range prime {
		prime[i] = true
	}
	prime[0], prime[1] = false, false
	//i := 2
	for i := int(math.Sqrt(float64(N))); i > 1; i-- {
		if prime[i] {
			for j := i * i; j <= N; j += i {
				prime[j] = false
			}
		}
	}
	var primeNbr []int
	for i := 2; i < len(prime); i++ {
		if prime[i] {
			primeNbr = append(primeNbr, i)
		}
	}
	return primeNbr
}
