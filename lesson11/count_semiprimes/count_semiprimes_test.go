package count_semiprimes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountSemiPrimes(t *testing.T) {
	{
		t.Run("Given P := []int{1, 4, 16},Q := []int{26, 10, 20},and N:=26 the function should return {10, 4, 0}", func(t *testing.T) {
			P := []int{1, 4, 16}
			Q := []int{26, 10, 20}
			N := 26

			expected := []int{10, 4, 0}

			got := Solution(N, P, Q)

			assert.Equal(t, expected, got)
		})
	}
}
