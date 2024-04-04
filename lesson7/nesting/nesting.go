package nesting

import "fmt"

// Detected time complexity:
// O(N)
// https://app.codility.com/demo/results/training34BJS8-QYX/
func Solution(S string) int {
	stack := make([]rune, 0)
	for _, v := range S {
		if v == '(' {
			stack = append(stack, v)
			fmt.Println(string(v))
		} else {
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			} else {
				return 0
			}
		}
	}
	if len(stack) == 0 {
		return 1
	}
	return 0
}
