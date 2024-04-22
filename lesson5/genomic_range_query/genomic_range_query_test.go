package genomic_range_query

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenomicRangeQuery(t *testing.T) {
	t.Run("Given P := []int{2, 5, 0}, Q := []int{4, 5, 6} and S := 'CAGCCTA', the function should return {2, 4, 1}", func(t *testing.T) {
		P := []int{2, 5, 0}
		Q := []int{4, 5, 6}
		S := "CAGCCTA"
		expected := []int{2, 4, 1}

		got := Solution(S, P, Q)

		assert.Equal(t, expected, got)
	})
	t.Run("Given P := []int{0}, Q := []int{0} and S := 'A', the function should return {1}", func(t *testing.T) {
		P := []int{0}
		Q := []int{0}
		S := "A"
		expected := []int{1}

		got := Solution(S, P, Q)

		assert.Equal(t, expected, got)
	})
	t.Run("Given P := []int{0, 0, 1}, Q := []int{0, 1, 1} and S := 'AC', the function should return {1, 1, 2}", func(t *testing.T) {
		P := []int{0, 0, 1}
		Q := []int{0, 1, 1}
		S := "AC"
		expected := []int{1, 1, 2}

		got := Solution(S, P, Q)

		assert.Equal(t, expected, got)
	})
}
