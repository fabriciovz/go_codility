package binary_gap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryGap(t *testing.T) {
	t.Run("Given N := 9, the function should return 2", func(t *testing.T) {
		N := 9
		expected := 2

		got := Solution(N)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N := 529, the function should return 4", func(t *testing.T) {
		N := 529
		expected := 4

		got := Solution(N)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N := 20, the function should return 1", func(t *testing.T) {
		N := 20
		expected := 1

		got := Solution(N)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N := 15, the function should return 0", func(t *testing.T) {
		N := 15
		expected := 0

		got := Solution(N)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N := 32, the function should return 0", func(t *testing.T) {
		N := 32
		expected := 0

		got := Solution(N)

		assert.Equal(t, expected, got)
	})

}
