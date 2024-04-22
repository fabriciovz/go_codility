package count_distinct_slices

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountDistinctSlices(t *testing.T) {
	t.Run("Given A := []int{3, 4, 5, 5, 2} and M := 6, then the function should return 9", func(t *testing.T) {
		A := []int{3, 4, 5, 5, 2}
		M := 6
		expected := 9

		got := Solution(M, A)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{1, 2, 3, 4, 5} and M := 6, then the function should return 15", func(t *testing.T) {
		A := []int{1, 2, 3, 4, 5}
		M := 6
		expected := 15

		got := Solution(M, A)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{1, 2, 3, 4, 7} and M := 6, then the function should return -1", func(t *testing.T) {
		A := []int{1, 2, 3, 4, 7}
		M := 6
		expected := -1

		got := Solution(M, A)

		assert.Equal(t, expected, got)
	})
}
