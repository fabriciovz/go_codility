package nesting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNesting(t *testing.T) {
	t.Run("Given A := []int{10, 2, 5, 1, 8, 20}, the function should return 1", func(t *testing.T) {
		S := "(()(())())"
		expected := 1

		got := Solution(S)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{10, 2, 5, 1, 8, 20}, the function should return 1", func(t *testing.T) {
		S := ""
		expected := 1

		got := Solution(S)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{10, 2, 5, 1, 8, 20}, the function should return 1", func(t *testing.T) {
		S := "())"
		expected := 0

		got := Solution(S)

		assert.Equal(t, expected, got)
	})
}
