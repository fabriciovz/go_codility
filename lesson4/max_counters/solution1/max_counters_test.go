package solution1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxCounters(t *testing.T) {
	t.Run("Given A={3, 4, 4, 6, 1, 4, 4}, the function should return {3, 2, 2, 4, 2}", func(t *testing.T) {
		A := []int{3, 4, 4, 6, 1, 4, 4}
		N := 5
		expected := []int{3, 2, 2, 4, 2}

		got := Solution(N, A)

		assert.Equal(t, expected, got)
	})
}
