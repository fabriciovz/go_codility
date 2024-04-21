package solution1

// Detected time complexity:
// O(N * log(log(N)))
// 100% score https://app.codility.com/demo/results/trainingQWZY62-S6P/
func Solution(A []int) int {
	N := len(A)
	if N < 3 {
		return 0
	}
	peaks := make([]int, 0)

	for i := 1; i < N-1; i++ {
		if A[i] > A[i-1] && A[i] > A[i+1] {
			peaks = append(peaks, i)
		}
	}
	numberOfPeaks := len(peaks)

	if numberOfPeaks == 0 {
		return 0
	}
	// Max number of slices can be N/numberOfPeaks
	maxNumberOfSlices := N / numberOfPeaks

	for blocksLenght := maxNumberOfSlices; blocksLenght <= N; blocksLenght++ {
		if N%blocksLenght != 0 {
			continue
		}
		groups := N / blocksLenght
		prevFound := -1
		count := 0
		for _, peak := range peaks {
			found := peak / blocksLenght
			if prevFound < found {
				// Peak found in the slice
				prevFound = found
				count++
			} else if prevFound > found {
				// Peak not found in the slice
				break
			} else if prevFound == found {
				// Another peak found in the same slice
				continue
			}
		}
		if groups == count {
			// If all slices contain at least one peak
			return groups
		}
	}
	return 0
}
