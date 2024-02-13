package go_coding

import (
	"strconv"
)

func Solution(n int) int {
	maxGap := 0
	currentGap := 0
	binary := strconv.FormatInt(int64(n), 2) // 1111011
	for _, digit := range binary {
		if digit == '0' {
			//also you can use: if string(v) == "0" {
			currentGap++
		} else {
			maxGap = Max(currentGap, maxGap)
			currentGap = 0
		}
	}
	return maxGap
}

func Max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
