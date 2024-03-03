package fish

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTriangle(t *testing.T) {
	t.Run("Given A := []int{10, 2, 5, 1, 8, 20}, the function should return 1", func(t *testing.T) {
		A := []int{4, 3, 2, 1, 5}
		B := []int{0, 1, 0, 0, 0}

		expected := 2

		got := Solution(A, B)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{10, 2, 5, 1, 8, 20}, the function should return 1", func(t *testing.T) {
		A := []int{4, 3, 2, 1, 1}
		B := []int{0, 1, 0, 0, 0}

		expected := 0

		got := Solution(A, B)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{10, 2, 5, 1, 8, 20}, the function should return 1", func(t *testing.T) {
		A := []int{4, 3, 2, 1, 1}
		B := []int{5, 1, 0, 0, 0}

		expected := 0

		got := Solution(A, B)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{10, 2, 5, 1, 8, 20}, the function should return 1", func(t *testing.T) {
		A := []int{4, 3, 2, 1, 5}
		B := []int{1, 0, 1, 0, 1}

		expected := 3

		got := Solution(A, B)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{10, 2, 5, 1, 8, 20}, the function should return 1", func(t *testing.T) {
		A := []int{4, 3, 2, 0, 5}
		B := []int{0, 1, 0, 0, 0}

		expected := 2

		got := Solution(A, B)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{10, 2, 5, 1, 8, 20}, the function should return 1", func(t *testing.T) {
		A := []int{4, 3, 2, 1, 5}
		B := []int{0, 1, 0, 0, 0}

		expected := 2

		got := Solution(A, B)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{10, 2, 5, 1, 8, 20}, the function should return 1", func(t *testing.T) {
		A := []int{4, 3, 2, 1, 5}
		B := []int{0, 1, 1, 0, 0}

		expected := 2

		got := Solution(A, B)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{10, 2, 5, 1, 8, 20}, the function should return 1", func(t *testing.T) {
		A := []int{4, 3, 2, 5, 6}
		B := []int{1, 0, 1, 0, 1}

		expected := 2

		got := Solution(A, B)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{10, 2, 5, 1, 8, 20}, the function should return 1", func(t *testing.T) {
		A := []int{7, 4, 3, 2, 5, 6}
		B := []int{0, 1, 1, 1, 0, 1}

		expected := 3

		got := Solution(A, B)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{10, 2, 5, 1, 8, 20}, the function should return 1", func(t *testing.T) {
		A := []int{3, 4, 2, 1, 5}
		B := []int{1, 0, 0, 0, 0}

		expected := 4

		got := Solution(A, B)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{10, 2, 5, 1, 8, 20}, the function should return 1", func(t *testing.T) {
		A := []int{3}
		B := []int{1}

		expected := 1

		got := Solution(A, B)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{10, 2, 5, 1, 8, 20}, the function should return 1", func(t *testing.T) {
		A := []int{3}
		B := []int{0}

		expected := 1

		got := Solution(A, B)

		assert.Equal(t, expected, got)
	})
}
