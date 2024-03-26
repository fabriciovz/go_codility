package genomic_range_query

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountDiv(t *testing.T) {
	t.Run("Given A := []int{3, 4, 4, 6, 1, 4, 4}, the function should return []int{3, 2, 2, 4, 2}", func(t *testing.T) {
		P := []int{2, 5, 0}
		Q := []int{4, 5, 6}
		S := "CAGCCTA"
		expected := []int{2, 4, 1}

		got := Solution(S, P, Q)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{3, 4, 4, 6, 1, 4, 4}, the function should return []int{3, 2, 2, 4, 2}", func(t *testing.T) {
		P := []int{0}
		Q := []int{0}
		S := "A"
		expected := []int{1}

		got := Solution(S, P, Q)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{3, 4, 4, 6, 1, 4, 4}, the function should return []int{3, 2, 2, 4, 2}", func(t *testing.T) {
		P := []int{0, 0, 1}
		Q := []int{0, 1, 1}
		S := "AC"
		expected := []int{1, 1, 2}

		got := Solution(S, P, Q)

		assert.Equal(t, expected, got)
	})
}
