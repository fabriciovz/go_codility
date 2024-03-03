package passing_cars

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPassingCars(t *testing.T) {
	t.Run("Given A={0,1,0,1,1}, then the function should return 5", func(t *testing.T) {
		a := []int{0, 1, 0, 1, 1}
		expected := 5

		got := Solution(a)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A={0, 1, 0, 1, 0, 1, 0, 1, 0, 1}, then the function should return 15", func(t *testing.T) {
		a := []int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1}
		expected := 15

		got := Solution(a)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A={0, 1, 1, 1, 0, 1, 1, 1, 0, 0, 1 }, then the function should return 13", func(t *testing.T) {
		a := []int{0, 1, 1, 1, 0, 1, 1, 1, 0, 0, 1}
		expected := 13

		got := Solution(a)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A={1, 1, 0, 1}, then the function should return 1", func(t *testing.T) {
		a := []int{1, 1, 0, 1}
		expected := 1

		got := Solution(a)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A={1, 1, 0}, then the function should return 0", func(t *testing.T) {
		a := []int{1, 1, 0}
		expected := 0

		got := Solution(a)

		assert.Equal(t, expected, got)
	})
}
