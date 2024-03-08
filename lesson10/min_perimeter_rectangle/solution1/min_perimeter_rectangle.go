package solution1

// 80% https://app.codility.com/demo/results/trainingP57CAA-79C/
const (
	MAX_INT = 1000000000
)

func Solution(N int) int {
	A := 1
	minPerimeter := MAX_INT

	if N < 1 || N > MAX_INT {
		return 0
	}
	for N/A >= 1 {
		B := N / A
		currentPerimeter := 0
		if N%A == 0 {
			currentPerimeter = 2 * (A + B)
			if currentPerimeter < minPerimeter {
				minPerimeter = currentPerimeter

			}
		}
		if B == A+1 {
			break
		}
		A++
	}
	return minPerimeter
}
