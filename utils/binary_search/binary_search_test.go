package binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	t.Run("Given N=0, then the function should return 0", func(t *testing.T) {
		N := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		X := 7
		expected := 7

		got := Solution(N, X)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N=0, then the function should return 0", func(t *testing.T) {
		N := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		X := 3
		expected := 3

		got := Solution(N, X)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N=0, then the function should return 0", func(t *testing.T) {
		N := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		X := 4
		expected := 4

		got := Solution(N, X)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N=0, then the function should return 0", func(t *testing.T) {
		N := []int{0, 1, 2, 3, 4, 5}
		X := 1
		expected := 1

		got := Solution(N, X)

		assert.Equal(t, expected, got)
	})
}
