package solution1

import "math"

// 66% score https://app.codility.com/demo/results/trainingW7Q4E2-TYP/
func Solution(A []int) int {
	N := len(A)

	if !IsvalidN(N) || !IsValidArray(A) {
		return 0
	}
	peaks := make([]int, 0)
	for i := 1; i < N-1; i++ {
		if A[i] > A[i-1] && A[i] > A[i+1] {
			peaks = append(peaks, i)
		}
	}
	NofPeaks := len(peaks)
	if NofPeaks <= 1 {
		return NofPeaks
	}
	sqrtOfN := int(math.Sqrt(float64(N))) + 1
	m := 0
	c := 1
	minPeak := Min(NofPeaks, sqrtOfN)
	for k := minPeak; k >= 1; k-- {
		lastF := 0
		c = 1
		for i := 1; i < NofPeaks; i++ {
			if peaks[i]-peaks[lastF] >= k && c < k {
				c++
				lastF = i
			}
		}
		if c < m {
			return m
		} else if m < c {
			m = c
		}
	}
	return m
}

func IsvalidN(N int) bool {
	if N < 1 || N > 400000 {
		return false
	}
	return true
}

func IsValidArray(A []int) bool {
	for i := 0; i < len(A); i++ {
		if A[i] < 0 || A[i] > 1000000000 {
			return false
		}
	}
	return true
}

func Min(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
