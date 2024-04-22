package count_semiprimes

// https://app.codility.com/demo/results/training7TH7H2-26S/
func Solution(N int, P []int, Q []int) []int {
	if !IsvalidN(N) || !IsvalidPQ(N, P, Q) {
		return []int{}
	}
	prime := make([]bool, N+1)
	semiPrime := make([]bool, N+1)
	acumSemiPrime := make([]int, N+1)
	result := make([]int, len(P))

	for i := range prime {
		prime[i] = true
	}
	prime[0], prime[1] = false, false
	//getting all prime numbers up to N
	i := 2
	for i*i <= N {
		if prime[i] {
			for j := i * i; j <= N; j += i {
				prime[j] = false
			}
		}
		i++
	}
	//filling semiprimes array
	for i = 0; i <= N; i++ {
		if prime[i] {
			for j := i; j <= N; j++ {
				if prime[j] && j*i <= N {
					currentSemiPrime := j * i
					semiPrime[currentSemiPrime] = true
				} else if j*i > N {
					break
				}
			}
		}
	}
	//using prefix sum acumulate semiprime numbers up to N
	for i = 2; i <= N; i++ {
		s := 0
		if semiPrime[i] {
			s++
		}
		acumSemiPrime[i] = acumSemiPrime[i-1] + s
	}
	//getting the semiprimes betwenn a range
	for i = 0; i < len(P); i++ {
		result[i] = acumSemiPrime[Q[i]] - acumSemiPrime[P[i]-1]
	}

	return result
}

func IsvalidN(N int) bool {
	if N < 1 || N > 50000 {
		return false
	}
	return true
}

func IsvalidPQ(N int, P []int, Q []int) bool {
	if len(P) != len(Q) {
		return false
	}
	if len(P) < 1 || len(P) > 30000 {
		return false
	}
	for i := 0; i < len(P); i++ {
		if P[i] > Q[i] {
			return false
		}
		if P[i] < 1 || P[i] > N {
			return false
		}
		if Q[i] < 1 || Q[i] > N {
			return false
		}
	}
	return true
}
