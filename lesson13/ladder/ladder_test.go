package ladder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLadder(t *testing.T) {
	t.Run("Given A := []int{4, 4, 5, 5, 1} and B := []int{3, 2, 4, 3, 1}, then the function should return {5, 1, 8, 0, 1}", func(t *testing.T) {
		A := []int{4, 4, 5, 5, 1}
		B := []int{3, 2, 4, 3, 1}

		expected := []int{5, 1, 8, 0, 1}

		got := Solution(A, B)

		assert.Equal(t, expected, got)
	})
}
