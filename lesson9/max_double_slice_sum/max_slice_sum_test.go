package max_double_slice_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxDoubleSliceSum(t *testing.T) {
	{
		t.Run("Given A = {3, 2, 6, -1, 4, 5, -1, 2}, the function should return 17", func(t *testing.T) {
			a := []int{3, 2, 6, -1, 4, 5, -1, 2}
			expected := 17

			got := Solution(a)

			assert.Equal(t, expected, got)
		})

		t.Run("Given A = {5, -7, 3, 5, -2, 4, -1}, the function should return 10", func(t *testing.T) {
			a := []int{5, -7, 3, 5, -2, 4, -1}
			expected := 10

			got := Solution(a)

			assert.Equal(t, expected, got)
		})
	}

}
