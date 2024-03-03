package solution1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFrogRiverOne(t *testing.T) {
	t.Run("Given A={1, 3, 1, 4, 2, 3, 5, 4}, the function should return 6", func(t *testing.T) {
		A := []int{3, 4, 4, 6, 1, 4, 4}
		N := 5
		expected := []int{3, 2, 2, 4, 2}

		got := Solution(N, A)

		assert.Equal(t, expected, got)
	})
}
