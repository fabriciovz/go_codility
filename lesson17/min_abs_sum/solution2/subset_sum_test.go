package solution2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinAbsSum(t *testing.T) {
	t.Run("Given A := []int{1, 5, 2, 10} and Target := 7, then the function should return true", func(t *testing.T) {
		A := []int{2, 4, 2, 3}
		expected := 1

		got := Solution(A)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{1, 5, 2, 10} and Target := 7, then the function should return true", func(t *testing.T) {
		A := []int{1, 5, 2, 2}
		expected := 0

		got := Solution(A)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{1, 5, 2, 10} and Target := 7, then the function should return true", func(t *testing.T) {
		A := []int{}
		expected := 0

		got := Solution(A)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{1, 5, 2, 10} and Target := 7, then the function should return true", func(t *testing.T) {
		A := []int{-100, 3, 2, 4}
		expected := 91

		got := Solution(A)

		assert.Equal(t, expected, got)
	})
}
