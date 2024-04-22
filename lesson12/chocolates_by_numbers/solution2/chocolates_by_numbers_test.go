package solution2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChocolatesByNumbers(t *testing.T) {
	t.Run("Given N := 10 and M := 4, then the function should return 5", func(t *testing.T) {
		N := 10
		M := 4
		expected := 5

		got := Solution(N, M)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N := 10 and M := 10, then the function should return 1", func(t *testing.T) {
		N := 10
		M := 10
		expected := 1

		got := Solution(N, M)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N := 10 and M := 1, then the function should return 10", func(t *testing.T) {
		N := 10
		M := 1
		expected := 10

		got := Solution(N, M)

		assert.Equal(t, expected, got)
	})

	t.Run("Given N := 10 and M := 9, then the function should return 10", func(t *testing.T) {
		N := 10
		M := 9
		expected := 10

		got := Solution(N, M)

		assert.Equal(t, expected, got)
	})
}
