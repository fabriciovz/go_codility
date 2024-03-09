package solution1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountFactor(t *testing.T) {
	{
		t.Run("Given N=24, the function should return {2, 3, 5, 7, 11, 13, 17, 19,23}", func(t *testing.T) {
			P := []int{1, 4, 16}
			Q := []int{26, 10, 20}
			N := 26

			expected := []int{10, 4, 0}

			got := Solution(N, P, Q)

			assert.Equal(t, expected, got)
		})
	}
}
