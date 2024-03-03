package solution1

// https://app.codility.com/demo/results/trainingAPW8G7-PT8/
func Solution(N int, M int) int {
	chocolatNumber := 0
	counter := 1
	for (chocolatNumber+M)%N != 0 {
		chocolatNumber = (chocolatNumber + M) % N
		counter++
	}
	return counter
}
