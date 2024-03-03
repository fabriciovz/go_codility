package solution2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountDiv(t *testing.T) {
	t.Run("Given A := []int{3, 4, 4, 6, 1, 4, 4}, the function should return []int{3, 2, 2, 4, 2}", func(t *testing.T) {
		A := 6
		B := 11
		K := 2
		expected := 3

		got := Solution(A, B, K)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{3, 4, 4, 6, 1, 4, 4}, the function should return []int{3, 2, 2, 4, 2}", func(t *testing.T) {
		A := 11
		B := 10
		K := 2
		expected := 0

		got := Solution(A, B, K)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{3, 4, 4, 6, 1, 4, 4}, the function should return []int{3, 2, 2, 4, 2}", func(t *testing.T) {
		A := 0
		B := 0
		K := 11
		expected := 1

		got := Solution(A, B, K)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{3, 4, 4, 6, 1, 4, 4}, the function should return []int{3, 2, 2, 4, 2}", func(t *testing.T) {
		A := 11
		B := 345
		K := 17
		expected := 20

		got := Solution(A, B, K)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{3, 4, 4, 6, 1, 4, 4}, the function should return []int{3, 2, 2, 4, 2}", func(t *testing.T) {
		A := 11
		B := 11
		K := 11
		expected := 1

		got := Solution(A, B, K)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{3, 4, 4, 6, 1, 4, 4}, the function should return []int{3, 2, 2, 4, 2}", func(t *testing.T) {
		A := 0
		B := 2000000000
		K := 1
		expected := 2000000001

		got := Solution(A, B, K)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{3, 4, 4, 6, 1, 4, 4}, the function should return []int{3, 2, 2, 4, 2}", func(t *testing.T) {
		A := 25
		B := 70
		K := 10
		expected := 5

		got := Solution(A, B, K)

		assert.Equal(t, expected, got)
	})
}
