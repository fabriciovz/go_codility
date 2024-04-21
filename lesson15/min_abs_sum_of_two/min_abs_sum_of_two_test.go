package min_abs_sum_of_two

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinAbsSumOfTwo(t *testing.T) {
	t.Run("Given A := []int{1, 4, -3}, the function should return 1", func(t *testing.T) {
		A := []int{1, 4, -3}
		expected := 1

		got := Solution(A)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{-8, 4, 5, -10, 3}, the function should return 3", func(t *testing.T) {
		A := []int{-8, 4, 5, -10, 3}
		expected := 3

		got := Solution(A)

		assert.Equal(t, expected, got)
	})
}
