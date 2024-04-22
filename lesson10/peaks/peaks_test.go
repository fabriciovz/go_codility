package solution1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPeaks(t *testing.T) {
	{
		t.Run("Given N := []int{1, 2, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2}, the function should return 3", func(t *testing.T) {
			N := []int{1, 2, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2}
			expected := 3

			got := Solution(N)

			assert.Equal(t, expected, got)
		})
		t.Run("Given N := []int{1, 2, 1, 2, 1, 1, 1, 1, 1, 1, 2, 1}, the function should return 2", func(t *testing.T) {
			N := []int{1, 2, 1, 2, 1, 1, 1, 1, 1, 1, 2, 1}
			expected := 2

			got := Solution(N)

			assert.Equal(t, expected, got)
		})
	}
}
