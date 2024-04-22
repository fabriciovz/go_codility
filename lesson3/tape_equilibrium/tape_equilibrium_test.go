package tape_equilibrium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTapeEquilibrium(t *testing.T) {
	t.Run("Given A = {3, 1, 2,4,3}, the the function should return 1", func(t *testing.T) {
		a := []int{3, 1, 2, 4, 3}
		expected := 1
		got := Solution(a)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A = {3}, the the function should return 0", func(t *testing.T) {
		a := []int{3}
		expected := 0
		got := Solution(a)

		assert.Equal(t, expected, got)

	})
	t.Run("Given A = {-30000, 1, 2,4,3}, the the function should return 0", func(t *testing.T) {
		a := []int{-30000, 1, 2, 4, 3}
		expected := 0
		got := Solution(a)

		assert.Equal(t, expected, got)

	})
	t.Run("Given A = {-1000, 1000}, the the function should return 2000", func(t *testing.T) {
		a := []int{-1000, 1000}
		expected := 2000
		got := Solution(a)

		assert.Equal(t, expected, got)

	})
}
