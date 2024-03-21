package tie_ropes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxNonOverlapping(t *testing.T) {
	t.Run("Given A := []int{1, 3, 7, 9, 9} and B := []int{5, 6, 8, 9, 10}, then the function should return 3", func(t *testing.T) {
		ropes := []int{1, 2, 3, 4, 1, 1, 3}
		K := 4
		expected := 3

		got := Solution(K, ropes)

		assert.Equal(t, expected, got)
	})
}
