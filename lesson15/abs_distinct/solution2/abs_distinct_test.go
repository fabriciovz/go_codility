package solution2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbsDistinct(t *testing.T) {
	t.Run("Given N := []int{-5, -3, -1, 0, 3, 6}, then the function should return 5", func(t *testing.T) {
		N := []int{-5, -3, -1, 0, 3, 6}
		expected := 5

		got := Solution(N)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N := []int{-10}, then the function should return 1", func(t *testing.T) {
		N := []int{-10}
		expected := 1

		got := Solution(N)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N := []int{-2, -2}, then the function should return 1", func(t *testing.T) {
		N := []int{-2, -2}
		expected := 1

		got := Solution(N)

		assert.Equal(t, expected, got)
	})
}
