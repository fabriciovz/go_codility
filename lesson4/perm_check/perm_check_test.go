package perm_check

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermCheck(t *testing.T) {
	t.Run("Given N := []int{4, 1, 3, 2}, then the function should return 1", func(t *testing.T) {
		N := []int{4, 1, 3, 2}
		expected := 1

		got := Solution(N)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N := []int{4, 1, 3}, then the function should return 0", func(t *testing.T) {
		N := []int{4, 1, 3}
		expected := 0

		got := Solution(N)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N := []int{1, 1}, then the function should return 0", func(t *testing.T) {
		N := []int{1, 1}
		expected := 0

		got := Solution(N)

		assert.Equal(t, expected, got)
	})

	t.Run("Given N := []int{2}, then the function should return 0", func(t *testing.T) {
		N := []int{2}
		expected := 0

		got := Solution(N)

		assert.Equal(t, expected, got)
	})
}
