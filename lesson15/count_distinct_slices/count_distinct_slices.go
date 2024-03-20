package count_distinct_slices

import "fmt"

// Detected time complexity:
// O(N)
// https://app.codility.com/demo/results/trainingBZVZ55-GJK/
func Solution(M int, A []int) int {
	if !IsValidInput(M, A) {
		return -1
	}
	N := len(A)
	front := 0
	caterpillar := map[int]int{}
	counter := 0

	for back := 0; back < N; back++ {
		for front < N && !Contains(caterpillar, A[front]) {
			//move head forward
			caterpillar[A[front]]++
			front++
			counter += len(caterpillar)
			if counter > 1000000000 {
				return 1000000000
			}
			PrintCaterpillar(caterpillar)
		}
		//move tail fordward
		delete(caterpillar, A[back])
		PrintCaterpillar(caterpillar)
	}
	return counter
}

func PrintCaterpillar(caterpillar map[int]int) {
	s := ""
	for i := range caterpillar {
		s += fmt.Sprintf("%d : ", i)
	}
	fmt.Println(s)
}

func Contains(caterpillar map[int]int, key int) bool {
	_, ok := caterpillar[key]
	if !ok {
		return false
	}
	return true
}

func IsValidInput(M int, A []int) bool {
	lenA := len(A)

	if lenA < 1 || lenA > 100000 {
		return false
	}
	if M < 0 || M > 100000 {
		return false
	}
	for i := 0; i < lenA; i++ {
		if A[i] < 0 || A[i] > M {
			return false
		}
	}
	return true
}
