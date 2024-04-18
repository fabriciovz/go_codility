package solution3

// Detected time complexity:
// O(N)
// 100% score https://app.codility.com/demo/results/trainingWDBAWA-BZG/
// Based on: https://codility.com/media/train/solution-flags.pdf
func Solution(A []int) int {
	N := len(A)
	peaks := make([]bool, N)
	peakCounter := 0
	for i := 1; i < N-1; i++ {
		if A[i] > A[i-1] && A[i] > A[i+1] {
			peaks[i] = true
			peakCounter++
		}
	}
	if peakCounter <= 2 {
		return peakCounter
	}
	next := make([]int, N)
	for i := N - 2; i >= 0; i-- {
		if peaks[i] {
			next[i] = i
		} else {
			next[i] = next[i+1]
		}
	}

	result := 0
	i := 1
	for (i-1)*i <= N {
		curFlag := 0
		curPos := 0
		for curPos < N && curFlag < i {
			curPos = next[curPos]
			//check for final value of next array
			if curPos == 0 {
				break
			}
			curFlag++
			curPos += i
		}

		if result < curFlag {
			result = curFlag
		}
		i++
	}
	return result
}
