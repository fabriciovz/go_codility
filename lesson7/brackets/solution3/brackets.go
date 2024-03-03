package solution3

// Detected time complexity:
// O(N)
func Solution(S string) int {
	stack := []rune{}
	mapChars := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	if !IsvalidN(len(S)) {
		return 0
	}
	// loop over string
	for _, r := range S {
		// if the current character is in the mapChars it means that is an open character
		// so push it its opposite in the stack
		if closer, ok := mapChars[r]; ok {
			stack = append(stack, closer)
			continue
		}
		// if it is not in the mapchars means that it is a closer
		// if the stack is empty then it is not a valid string
		// if the top char is not the same as closer means that it is invalid
		l := len(stack) - 1
		if l < 0 || r != stack[l] {
			return 0
		}
		// take the last element off the stack
		stack = stack[:l]
	}
	if len(stack) == 0 {
		return 1
	}
	return 0
}

func IsvalidN(N int) bool {
	if N < 0 || N > 200000 {
		return false
	}
	return true
}
