package max_non_overlapping_segments

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxNonOverlapping(t *testing.T) {
	t.Run("Given A := []int{1, 3, 7, 9, 9} and B := []int{5, 6, 8, 9, 10}, then the function should return 3", func(t *testing.T) {
		A := []int{1, 3, 7, 9, 9}
		B := []int{5, 6, 8, 9, 10}
		expected := 3

		got := Solution(A, B)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{1, 4, 7, 9} and B := []int{3, 6, 8, 10}, then the function should return 4", func(t *testing.T) {
		A := []int{1, 4, 7, 9}
		B := []int{3, 6, 8, 10}
		expected := 4

		got := Solution(A, B)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{} and B := []int{}, then the function should return 0", func(t *testing.T) {
		A := []int{}
		B := []int{}
		expected := 0

		got := Solution(A, B)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{1} and B := []int{2}, then the function should return 1", func(t *testing.T) {
		A := []int{1}
		B := []int{2}
		expected := 1

		got := Solution(A, B)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{1, 2, 3, 1} and B := []int{2, 3, 4, 3}, then the function should return 2", func(t *testing.T) {
		A := []int{1, 2, 3, 1}
		B := []int{2, 3, 4, 3}
		expected := 2

		got := Solution(A, B)

		assert.Equal(t, expected, got)
	})
}
