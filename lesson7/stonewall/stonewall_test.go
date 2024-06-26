package stonewall

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStoneWall(t *testing.T) {
	t.Run("Given A := []int{8, 8, 5, 7, 9, 8, 7, 4, 8}, the function should return 7", func(t *testing.T) {
		H := []int{8, 8, 5, 7, 9, 8, 7, 4, 8}
		expected := 7

		got := Solution(H)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{1, 1, 2, 2, 3, 3, 2, 2, 1, 1}, the function should return 3", func(t *testing.T) {
		H := []int{1, 1, 2, 2, 3, 3, 2, 2, 1, 1}
		expected := 3

		got := Solution(H)

		assert.Equal(t, expected, got)
	})
}
