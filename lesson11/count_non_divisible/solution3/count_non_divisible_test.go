package solution3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountFactor(t *testing.T) {
	{
		t.Run("Given A := []int{3, 1, 2, 3, 6}, the function should return {2, 4, 3, 2, 0}", func(t *testing.T) {
			a := []int{3, 1, 2, 3, 6}
			expected := []int{2, 4, 3, 2, 0}

			got := Solution(a)

			assert.Equal(t, expected, got)
		})
		t.Run("Given A := []int{2}, the function should return {0}", func(t *testing.T) {
			a := []int{2}
			expected := []int{0}

			got := Solution(a)

			assert.Equal(t, expected, got)
		})
	}
}
