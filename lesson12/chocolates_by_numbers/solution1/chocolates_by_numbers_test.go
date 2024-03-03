package solution1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEuclideanAlgorithm(t *testing.T) {
	t.Run("Given A=10 and B=4, then the function should return 5", func(t *testing.T) {
		N := 10
		M := 4
		expected := 5

		got := Solution(N, M)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A=10 and B=10, then the function should return 1", func(t *testing.T) {
		N := 10
		M := 10
		expected := 1

		got := Solution(N, M)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A=10 and B=1, then the function should return 10", func(t *testing.T) {
		N := 10
		M := 1
		expected := 10

		got := Solution(N, M)

		assert.Equal(t, expected, got)
	})

	t.Run("Given A=10 and B=9, then the function should return 10", func(t *testing.T) {
		N := 10
		M := 9
		expected := 10

		got := Solution(N, M)

		assert.Equal(t, expected, got)
	})
}
