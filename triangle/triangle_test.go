package triangle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTriangle(t *testing.T) {
	t.Run("Given A := []int{10, 2, 5, 1, 8, 20}, the function should return 1", func(t *testing.T) {
		A := []int{10, 2, 5, 1, 8, 20}
		expected := 1

		got := Solution(A)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{10, 50, 5, 1}, the function should return 0", func(t *testing.T) {
		A := []int{10, 50, 5, 1}
		expected := 0

		got := Solution(A)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{1, 1, 1, 1, 5, 5, 5}, the function should return 1", func(t *testing.T) {
		A := []int{1, 1, 1, 1, 5, 5, 5}
		expected := 1

		got := Solution(A)

		assert.Equal(t, expected, got)
	})
}
