package solution2

import "fmt"

// Detected time complexity:
// O(N)
func Solution(S string) int {
	mapChars := map[byte]int{'{': 1,
		'}': -1,
		'(': 2,
		')': -2,
		'[': 3,
		']': -3,
	}
	if !IsValidArray(S, mapChars) || !IsvalidN(len(S)) {
		return 0
	}
	stackChars := &stack{}
	for i := 0; i < len(S); i++ {
		fmt.Println(string(S[i]))
		v, ok := mapChars[S[i]]
		if !ok {
			continue
		}
		if v > 0 {
			stackChars.Push(S[i])
		} else {
			if stackChars.IsEmpty() {
				return 0
			} else {
				if mapChars[stackChars.Top()]+mapChars[S[i]] == 0 {
					stackChars.Pop()
				}
			}
		}
		fmt.Println(v)
	}
	fmt.Println(stackChars)
	if stackChars.IsEmpty() {
		return 1
	}
	return 0
}

func IsValidArray(A string, validChars map[byte]int) bool {
	for i := range A {
		_, ok := validChars[A[i]]
		if !ok {
			return false
		}
	}
	return true
}

func IsvalidN(N int) bool {
	if N < 0 || N > 200000 {
		return false
	}
	return true
}
