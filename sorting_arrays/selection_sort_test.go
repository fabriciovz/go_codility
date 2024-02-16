package sorting_arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	t.Run("Given A = {8, 7, 9, 2, 3, 1, 5, 4, 2, 6}, then the function should return {1, 2, 2, 3, 4, 5, 6, 7, 8, 9}", func(t *testing.T) {
		a := []int{8, 7, 9, 2, 3, 1, 5, 4, 2, 6}
		expected := []int{1, 2, 2, 3, 4, 5, 6, 7, 8, 9}

		got := SelectionSort(a)

		assert.Equal(t, expected, got)

	})
}
