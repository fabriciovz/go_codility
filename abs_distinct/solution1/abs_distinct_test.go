package solution1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbsDistinct(t *testing.T) {
	t.Run("Given N=0, then the function should return 0", func(t *testing.T) {
		N := []int{-5, -3, -1, 0, 3, 6}
		expected := 5

		got := Solution(N)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N=0, then the function should return 0", func(t *testing.T) {
		N := []int{-10}
		expected := 1

		got := Solution(N)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N=0, then the function should return 0", func(t *testing.T) {
		N := []int{-2, -2}
		expected := 1

		got := Solution(N)

		assert.Equal(t, expected, got)
	})
}
