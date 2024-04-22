package solution1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOddOccurrences(t *testing.T) {
	t.Run("Given A ={9, 3, 9}, the function should return 3", func(t *testing.T) {
		a := []int{9, 3, 9}
		expected := 3

		got := Solution(a)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A ={9, 3, 9, 3, 9, 7, 9}, the function should return 7", func(t *testing.T) {
		a := []int{9, 3, 9, 3, 9, 7, 9}
		expected := 7

		got := Solution(a)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A ={9}, the function should return 9", func(t *testing.T) {
		a := []int{9}
		expected := 9

		got := Solution(a)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A ={7, 6, 5, 5, 1, 7, 6}, the function should return 1", func(t *testing.T) {
		a := []int{7, 6, 5, 5, 1, 7, 6}
		expected := 1

		got := Solution(a)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A ={2, 3, 5, 4, 5, 2, 4, 3, 5, 2, 4, 4, 2}, the function should return 5", func(t *testing.T) {
		a := []int{2, 3, 5, 4, 5, 2, 4, 3, 5, 2, 4, 4, 2}
		expected := 5

		got := Solution(a)

		assert.Equal(t, expected, got)
	})

}
