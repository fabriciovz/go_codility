package solution3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinAbsSum(t *testing.T) {
	t.Run("Given A := []int{1, 3, 7, 9, 9} and B := []int{5, 6, 8, 9, 10}, then the function should return 3", func(t *testing.T) {
		A := []int{1, 5, 2, -2}
		expected := 0

		got := Solution(A)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{1, 3, 7, 9, 9} and B := []int{5, 6, 8, 9, 10}, then the function should return 3", func(t *testing.T) {
		A := []int{}
		expected := 0

		got := Solution(A)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{1, 3, 7, 9, 9} and B := []int{5, 6, 8, 9, 10}, then the function should return 3", func(t *testing.T) {
		A := []int{-100, 3, 2, 4}
		expected := 91

		got := Solution(A)

		assert.Equal(t, expected, got)
	})
}
