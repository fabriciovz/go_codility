package solution1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNailingPlanks(t *testing.T) {
	t.Run("Given A := []int{4, 5, 9, 10},B := []int{1, 4, 5, 8} and  C := []int{4, 6, 7, 10, 2}, then the function should return -1", func(t *testing.T) {
		//arrange
		A := []int{4, 5, 9, 10}
		B := []int{1, 4, 5, 8}
		C := []int{4, 6, 7, 10, 2}
		expected := -1

		//act
		var got = Solution(A, B, C)

		//assert
		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{1, 2, 3},B := []int{4, 5} and  C := []int{4, 6, 7, 10, 2}, then the function should return -1", func(t *testing.T) {
		//arrange
		A := []int{1, 2, 3}
		B := []int{4, 5}
		C := []int{4, 6, 7, 10, 2}
		expected := -1
		//act
		var got = Solution(A, B, C)

		//assert
		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{0, 1, 1, 2},B := []int{1, 1, 1, 2} and  C := []int{1}, then the function should return -1", func(t *testing.T) {
		//arrange
		A := []int{0, 1, 1, 2}
		B := []int{1, 1, 1, 2}
		C := []int{1}
		expected := -1

		//act
		var got = Solution(A, B, C)

		//assert
		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{1, 1, 1, 1},B := []int{1, 1, 1, 3} and  C := []int{1}, then the function should return -1", func(t *testing.T) {
		//arrange
		A := []int{1, 1, 1, 1}
		B := []int{1, 1, 1, 3}
		C := []int{1}
		expected := -1

		//act
		var got = Solution(A, B, C)

		//assert
		assert.Equal(t, expected, got)
	})

	t.Run("Given A := []int{1, 4, 5, 8},B := []int{4, 5, 9, 10} and  C := []int{4, 6, 7, 10, 2}, then the function should return 4", func(t *testing.T) {
		//arrange
		A := []int{1, 4, 5, 8}
		B := []int{4, 5, 9, 10}
		C := []int{4, 6, 7, 10, 2}
		expected := 4

		//act
		var got = Solution(A, B, C)

		//assert
		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{2},B := []int{2} and  C := []int{1}, then the function should return -1", func(t *testing.T) {
		//arrange
		A := []int{2}
		B := []int{2}
		C := []int{1}
		expected := -1

		//act
		var got = Solution(A, B, C)

		//assert
		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{1, 4, 5, 8},B := []int{4, 5, 9, 10} and  C := []int{6, 4, 2, 1, 2}, then the function should return -1", func(t *testing.T) {
		//arrange
		A := []int{1, 4, 5, 8}
		B := []int{4, 5, 9, 10}
		C := []int{6, 4, 2, 1, 2}
		expected := -1

		//act
		var got = Solution(A, B, C)

		//assert
		assert.Equal(t, expected, got)
	})
}
