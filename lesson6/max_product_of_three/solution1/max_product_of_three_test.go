package solution1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProductOfThree(t *testing.T) {
	t.Run("Given A={1, 3, 1, 4, 2, 3, 5, 4}, the function should return 6", func(t *testing.T) {
		a := []int{-3, 1, 2, -2, 5, 6}
		expected := 60

		got := Solution(a)

		assert.Equal(t, expected, got)
	})

	t.Run("Given A={1, 3, 1, 4, 2, 3, 5, 4}, the function should return 6", func(t *testing.T) {
		a := []int{-10, -2, -4}
		expected := -80

		got := Solution(a)

		assert.Equal(t, expected, got)
	})

	t.Run("Given A={1, 3, 1, 4, 2, 3, 5, 4}, the function should return 6", func(t *testing.T) {
		a := []int{-10, -2, -4, 3}
		expected := 120

		got := Solution(a)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A={1, 3, 1, 4, 2, 3, 5, 4}, the function should return 6", func(t *testing.T) {
		a := []int{-5, -6, -4, -7, -10}
		expected := -120

		got := Solution(a)

		assert.Equal(t, expected, got)
	})

}
