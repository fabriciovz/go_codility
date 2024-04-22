package min_avg_two_slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountDiv(t *testing.T) {
	t.Run("Given A := []int{4, 2, 2, 5, 1, 5, 8}, the function should return 1", func(t *testing.T) {
		A := []int{4, 2, 2, 5, 1, 5, 8}
		expected := 1

		got := Solution(A)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{10000, -10000}, the function should return 0", func(t *testing.T) {
		A := []int{10000, -10000}
		expected := 0

		got := Solution(A)

		assert.Equal(t, expected, got)
	})
}
