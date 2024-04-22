package min_max_division

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinMaxDiv(t *testing.T) {
	t.Run("Given A := []int{2, 1, 5, 1, 2, 2, 2}, K := 3 and M := 5, then the function should return 6", func(t *testing.T) {
		A := []int{2, 1, 5, 1, 2, 2, 2}
		K := 3
		M := 5
		expected := 6

		got := Solution(K, M, A)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{2, 1, 5, 1, 2, 2, 2}, K := 7 and M := 5, then the function should return 5", func(t *testing.T) {
		A := []int{2, 1, 5, 1, 2, 2, 2}
		K := 7
		M := 5
		expected := 5

		got := Solution(K, M, A)

		assert.Equal(t, expected, got)
	})
}
