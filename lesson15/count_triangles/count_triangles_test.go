package count_triangles

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountTriangles(t *testing.T) {
	t.Run("Given A := []int{10, 2, 5, 1, 8, 12}, the function should return 4", func(t *testing.T) {
		A := []int{10, 2, 5, 1, 8, 12}
		expected := 4

		got := Solution(A)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{4, 6, 3, 7}, the function should return 3", func(t *testing.T) {
		A := []int{4, 6, 3, 7}
		expected := 3

		got := Solution(A)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{10, 2, 5, 1, 8, 12}, the function should return 4", func(t *testing.T) {
		A := []int{10, 21, 22, 100, 101, 200, 300}
		expected := 6

		got := Solution(A)

		assert.Equal(t, expected, got)
	})
}
