package common_prime_divisors

// 100% https://app.codility.com/demo/results/trainingY6ZNH7-EVH/
func Solution(A []int, B []int) int {
	N := len(A)
	counter := 0
	for i := 0; i < N; i++ {
		if CheckCommonPrimeDivisor(A[i], B[i]) {
			counter++
		}
	}
	return counter
}

/*
if all the prime factors of Y are factors of X,
then Y/gcd(X,Y) will be 1 or also contain only factors of X
*/
func CheckCommonPrimeDivisor(A int, B int) bool {
	if A == B {
		return true
	} else if A == 1 || B == 1 {
		return false
	} else {
		gcd := GCD(A, B, 1)
		if gcd == 1 {
			return false
		} else {
			for GCD(A, gcd, 1) != 1 {
				A /= GCD(A, gcd, 1)
			}
			for GCD(B, gcd, 1) != 1 {
				B /= GCD(B, gcd, 1)
			}
			if A == 1 && B == 1 {
				return true
			} else {
				return false
			}
		}
	}
}

func GCD(A, B, res int) int {
	if A == B {
		return res * A
	} else if A%2 == 0 && B%2 == 0 {
		return GCD(A/2, B/2, 2*res)
	} else if A%2 == 0 {
		return GCD(A/2, B, res)
	} else if B%2 == 0 {
		return GCD(A, B/2, res)
	} else if A > B {
		return GCD(A-B, B, res)
	} else {
		return GCD(A, B-A, res)
	}
}
