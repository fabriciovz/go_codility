package frog_river_one

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFrogRiverOne(t *testing.T) {
	t.Run("Given A={1, 3, 1, 4, 2, 3, 5, 4}, the function should return 6", func(t *testing.T) {
		a := []int{1, 3, 1, 4, 2, 3, 5, 4}
		x := 5
		expected := 6

		got := Solution(x, a)

		assert.Equal(t, expected, got)
	})

	t.Run("Given A={1}, the function should return 0", func(t *testing.T) {
		a := []int{1}
		x := 1
		expected := 0

		got := Solution(x, a)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A={2, 2, 2, 2, 2}, the function should return -1", func(t *testing.T) {
		a := []int{2, 2, 2, 2, 2}
		x := 2
		expected := -1

		got := Solution(x, a)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A={1, 3, 1, 3, 2, 1, 3}, the function should return -1", func(t *testing.T) {
		a := []int{1, 3, 1, 3, 2, 1, 3}
		x := 3
		expected := 4

		got := Solution(x, a)

		assert.Equal(t, expected, got)
	})
}
