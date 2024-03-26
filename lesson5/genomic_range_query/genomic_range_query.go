package genomic_range_query

//Detected time complexity:
//O(N + M)
//100% https://app.codility.com/demo/results/training97F54B-ZR3/

func Solution(S string, P []int, Q []int) []int {
	mapImpact := map[rune]int{
		'A': 1,
		'C': 2,
		'G': 3,
		'T': 4,
	}
	mapCounter := make(map[rune][]int, len(mapImpact))
	for key, _ := range mapImpact {
		mapCounter[key] = make([]int, len(S)+1)
	}
	a := 0
	c := 0
	g := 0
	t := 0

	for i := 1; i < len(S)+1; i++ {
		if S[i-1] == 'A' {
			a++
		} else if S[i-1] == 'C' {
			c++
		} else if S[i-1] == 'G' {
			g++
		} else if S[i-1] == 'T' {
			t++
		}
		mapCounter['A'][i] = a
		mapCounter['C'][i] = c
		mapCounter['G'][i] = g
		mapCounter['T'][i] = t
	}
	result := make([]int, 0)

	for i := 0; i < len(P); i++ {
		startIndex := P[i]
		endIndex := Q[i] + 1

		if mapCounter['A'][startIndex] != mapCounter['A'][endIndex] {
			result = append(result, mapImpact['A'])
		} else if mapCounter['C'][startIndex] != mapCounter['C'][endIndex] {
			result = append(result, mapImpact['C'])
		} else if mapCounter['G'][startIndex] != mapCounter['G'][endIndex] {
			result = append(result, mapImpact['G'])
		} else if mapCounter['T'][startIndex] != mapCounter['T'][endIndex] {
			result = append(result, mapImpact['T'])
		}
	}
	return result
}
