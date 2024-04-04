package sorting_arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountingSort(t *testing.T) {
	t.Run("Given A = {8, 7, 9, 2, 3, 1, 5, 4, 2, 6}, then the function should return {1, 2, 2, 3, 4, 5, 6, 7, 8, 9}", func(t *testing.T) {
		a := []int{8, 7, 9, 2, 3, 1, 5, 4, 2, 6}
		expected := []int{1, 2, 2, 3, 4, 5, 6, 7, 8, 9}

		got := CountingSort(a)

		assert.Equal(t, expected, got)

	})
	t.Run("Given A = {4,9,2,4,7,6}, then the function should return {2, 4, 4, 6, 7, 9}", func(t *testing.T) {
		a := []int{4, 9, 2, 4, 7, 6}
		expected := []int{2, 4, 4, 6, 7, 9}

		got := CountingSort(a)

		assert.Equal(t, expected, got)

	})
	t.Run("Given A = {4,9,2,4,7,6}, then the function should return {2, 4, 4, 6, 7, 9}", func(t *testing.T) {
		a := []int{1, 0, 3, 1, 3, 1}
		expected := []int{0, 1, 1, 1, 3, 3}

		got := CountingSort(a)

		assert.Equal(t, expected, got)

	})
}
