package solution1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEuclideanAlgorithm(t *testing.T) {
	t.Run("Given N=0, then the function should return 0", func(t *testing.T) {
		N := 0
		expected := 0

		got := Solution(N)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N=1, then the function should return 1", func(t *testing.T) {
		N := 1
		expected := 1

		got := Solution(N)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N=2, then the function should return 1", func(t *testing.T) {
		N := 2
		expected := 1

		got := Solution(N)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N=3, then the function should return 2", func(t *testing.T) {
		N := 3
		expected := 2

		got := Solution(N)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N=11, then the function should return 89", func(t *testing.T) {
		N := 11
		expected := 89

		got := Solution(N)

		assert.Equal(t, expected, got)
	})
}
