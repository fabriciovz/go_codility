package solution2

import "math"

// Detected time complexity:
// O(N)
// 100% score https://app.codility.com/demo/results/trainingUPFEKH-GQ9/
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
	maxFlags := int(math.Ceil(math.Sqrt(float64(peaks[NofPeaks-1] - peaks[0]))))

	for flags := maxFlags; flags > 1; flags-- {
		startIndex := 0
		endIndex := NofPeaks - 1

		startFlag := peaks[startIndex]
		endFlag := peaks[endIndex]

		flagsPlaced := 2

		for startIndex < endIndex {
			if flagsPlaced == flags {
				return flags
			}
			startIndex++
			endIndex--
			if peaks[startIndex] >= startFlag+flags {
				if peaks[startIndex] <= endFlag-flags {
					flagsPlaced++
					startFlag = peaks[startIndex]
				}
			}
			if peaks[endIndex] >= startFlag+flags {
				if peaks[endIndex] <= endFlag-flags {
					flagsPlaced++
					endFlag = peaks[endIndex]
				}
			}
		}
	}
	return 0
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
