package count_distinct_slices

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountDistinctSlices(t *testing.T) {
	t.Run("", func(t *testing.T) {
		A := []int{3, 4, 5, 5, 2}
		M := 6
		expected := 9

		got := Solution(M, A)

		assert.Equal(t, expected, got)
	})
	t.Run("", func(t *testing.T) {
		A := []int{1, 2, 3, 4, 5}
		M := 6
		expected := 15

		got := Solution(M, A)

		assert.Equal(t, expected, got)
	})
	t.Run("", func(t *testing.T) {
		A := []int{1, 2, 3, 4, 7}
		M := 6
		expected := -1

		got := Solution(M, A)

		assert.Equal(t, expected, got)
	})
	t.Run("", func(t *testing.T) {
		var A []int
		M := 100001
		for i := 0; i < 100001; i++ {
			A = append(A, i)
		}
		expected := -1

		got := Solution(M, A)

		assert.Equal(t, expected, got)
	})

}
