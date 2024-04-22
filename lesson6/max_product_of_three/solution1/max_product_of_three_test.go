package solution1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProductOfThree(t *testing.T) {
	t.Run("Given A={-3, 1, 2, -2, 5, 6}, the function should return 60", func(t *testing.T) {
		a := []int{-3, 1, 2, -2, 5, 6}
		expected := 60

		got := Solution(a)

		assert.Equal(t, expected, got)
	})

	t.Run("Given A={-10, -2, -4}, the function should return -80", func(t *testing.T) {
		a := []int{-10, -2, -4}
		expected := -80

		got := Solution(a)

		assert.Equal(t, expected, got)
	})

	t.Run("Given A={-10, -2, -4, 3}, the function should return 120", func(t *testing.T) {
		a := []int{-10, -2, -4, 3}
		expected := 120

		got := Solution(a)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A={-5, -6, -4, -7, -10}, the function should return -120", func(t *testing.T) {
		a := []int{-5, -6, -4, -7, -10}
		expected := -120

		got := Solution(a)

		assert.Equal(t, expected, got)
	})

}
