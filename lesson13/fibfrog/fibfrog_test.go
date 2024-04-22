package solution2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEuclideanAlgorithm(t *testing.T) {
	t.Run("Given N := []int{0, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0}, then the function should return 3", func(t *testing.T) {
		N := []int{0, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0}
		expected := 3

		got := Solution(N)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N := []int{}, then the function should return 1", func(t *testing.T) {
		var N []int
		expected := 1

		got := Solution(N)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N := []int{1}, then the function should return 1", func(t *testing.T) {
		N := []int{1}
		expected := 1

		got := Solution(N)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N := []int{1, 1}, then the function should return 1", func(t *testing.T) {
		N := []int{1, 1}
		expected := 1

		got := Solution(N)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N := []int{0, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0}, then the function should return 3", func(t *testing.T) {
		N := []int{0, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0}
		expected := 3

		got := Solution(N)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N := []int{0, 0, 0}, then the function should return -1", func(t *testing.T) {
		N := []int{0, 0, 0}
		expected := -1

		got := Solution(N)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N := []int{0, 0, 0, 0}, then the function should return 1", func(t *testing.T) {
		N := []int{0, 0, 0, 0}
		expected := 1

		got := Solution(N)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N := []int{1, 0, 0, 0, 0}, then the function should return 2", func(t *testing.T) {
		N := []int{1, 0, 0, 0, 0}
		expected := 2

		got := Solution(N)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N := []int{0, 1, 0, 1, 0}, then the function should return 3", func(t *testing.T) {
		N := []int{0, 1, 0, 1, 0}
		expected := 3

		got := Solution(N)

		assert.Equal(t, expected, got)
	})
}
