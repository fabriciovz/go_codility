package solution2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProductOfThree(t *testing.T) {
	t.Run("Given A={1, 5, 2, 1, 4, 0}, the function should return 11", func(t *testing.T) {
		A := []int{1, 5, 2, 1, 4, 0}
		expected := 11

		got := Solution(A)

		assert.Equal(t, expected, got)
	})
}
