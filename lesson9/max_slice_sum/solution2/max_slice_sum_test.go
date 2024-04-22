package solution2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSliceSum(t *testing.T) {
	{
		t.Run("Given A = {-2, 1, -3, 4, -1, 2, 1, -5, 4}, the function should return 6", func(t *testing.T) {
			a := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
			expected := 6

			got := Solution(a)

			assert.Equal(t, expected, got)
		})
		t.Run("Given A = {5, -7, 3, 5, -2, 4, -1}, the function should return 10", func(t *testing.T) {
			a := []int{5, -7, 3, 5, -2, 4, -1}
			expected := 10

			got := Solution(a)

			assert.Equal(t, expected, got)
		})
		t.Run("Given A = {3, 2, -6, 4, 0}, the function should return 5", func(t *testing.T) {
			a := []int{3, 2, -6, 4, 0}
			expected := 5

			got := Solution(a)

			assert.Equal(t, expected, got)
		})
		t.Run("Given A = {3}, the function should return 3", func(t *testing.T) {
			a := []int{3}
			expected := 3

			got := Solution(a)

			assert.Equal(t, expected, got)
		})
		t.Run("Given A = {-2, 1}, the function should return 1", func(t *testing.T) {
			a := []int{-2, 1}
			expected := 1

			got := Solution(a)

			assert.Equal(t, expected, got)
		})
		t.Run("Given A = {-2, -2}, the function should return -2", func(t *testing.T) {
			a := []int{-2, -2}
			expected := -2

			got := Solution(a)

			assert.Equal(t, expected, got)
		})
		t.Run("Given A = {-2, -2}, the function should return 0", func(t *testing.T) {
			var a []int
			expected := 0

			got := Solution(a)

			assert.Equal(t, expected, got)
		})
		t.Run("Given A = {-2, -3, -6, -12, -1, -52}, the function should return -1", func(t *testing.T) {
			a := []int{-2, -3, -6, -12, -1, -52}
			expected := -1

			got := Solution(a)

			assert.Equal(t, expected, got)
		})
		t.Run("Given A = {2, 3, 6, 12, 1, 52}, the function should return 76", func(t *testing.T) {
			a := []int{2, 3, 6, 12, 1, 52}
			expected := 76

			got := Solution(a)

			assert.Equal(t, expected, got)
		})
	}

}
