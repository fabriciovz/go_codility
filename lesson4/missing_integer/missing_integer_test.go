package missing_integer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFrogRiverOne(t *testing.T) {
	t.Run("Given A={1, 3, 1, 4, 2, 3, 5, 4}, the function should return 6", func(t *testing.T) {
		A := []int{1, 3, 6, 4, 1, 2}
		expected := 5

		got := Solution(A)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A={1, 3, 1, 4, 2, 3, 5, 4}, the function should return 6", func(t *testing.T) {
		A := []int{1, 2, 3}
		expected := 4

		got := Solution(A)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A={1, 3, 1, 4, 2, 3, 5, 4}, the function should return 6", func(t *testing.T) {
		A := []int{-1, -3}
		expected := 1

		got := Solution(A)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A={1, 3, 1, 4, 2, 3, 5, 4}, the function should return 6", func(t *testing.T) {
		A := []int{-1000000, 1000000}
		expected := 1

		got := Solution(A)

		assert.Equal(t, expected, got)
	})
}
