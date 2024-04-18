package solution3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlags(t *testing.T) {
	{
		t.Run("Given N := []int{1, 5, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2}, the function should return 3", func(t *testing.T) {
			N := []int{1, 5, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2}
			expected := 3

			got := Solution(N)

			assert.Equal(t, expected, got)
		})
		t.Run("Given N := []int{2, 6, 4, 3, 2, 1, 4, 3, 4, 3, 5, 1}, the function should return 3", func(t *testing.T) {
			N := []int{2, 6, 4, 3, 2, 1, 4, 3, 4, 3, 5, 1}
			expected := 3

			got := Solution(N)

			assert.Equal(t, expected, got)
		})
		t.Run("Given N := []int{2, 6, 4, 3, 2, 1, 4, 3, 4, 3, 5, 1}, the function should return 3", func(t *testing.T) {
			N := []int{1, 2, 3, 4, 3, 2, 1, 3, 5, 4, 1}
			expected := 2

			got := Solution(N)

			assert.Equal(t, expected, got)
		})
		t.Run("Given N := []int{1, 2, 1, 2, 1, 2, 1, 2, 1}, the function should return 2", func(t *testing.T) {
			N := []int{1, 2, 1, 2, 1, 2, 1, 2, 1, 2}
			expected := 2

			got := Solution(N)

			assert.Equal(t, expected, got)
		})
		t.Run("Given N := []int{5}, the function should return 0", func(t *testing.T) {
			N := []int{5}
			expected := 0

			got := Solution(N)

			assert.Equal(t, expected, got)
		})
	}
}
