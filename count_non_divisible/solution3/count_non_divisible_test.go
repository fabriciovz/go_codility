package solution3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountFactor(t *testing.T) {
	{
		t.Run("Given N=24, the function should return {2, 3, 5, 7, 11, 13, 17, 19,23}", func(t *testing.T) {
			a := []int{3, 1, 2, 3, 6}
			expected := []int{2, 4, 3, 2, 0}

			got := Solution(a)

			assert.Equal(t, expected, got)
		})
		t.Run("Given N=24, the function should return {2, 3, 5, 7, 11, 13, 17, 19,23}", func(t *testing.T) {
			a := []int{2}
			expected := []int{0}

			got := Solution(a)

			assert.Equal(t, expected, got)
		})
	}
}
