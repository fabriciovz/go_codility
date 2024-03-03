package sieve_of_eratosthenes

// Find all prime numbers till N
func Solution(N int) []int {
	prime := make([]bool, N+1)
	for i := range prime {
		prime[i] = true
	}
	i := 2
	for i*i < N {
		if prime[i] {
			for j := i * i; j <= N; j += i {
				prime[j] = false
			}
		}
		i++
	}
	var primeNbr []int
	for i = 2; i < len(prime); i++ {
		if prime[i] {
			primeNbr = append(primeNbr, i)
		}
	}
	return primeNbr
}
