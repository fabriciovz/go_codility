package distinct

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDistinct(t *testing.T) {
	t.Run("Given A = {8, 7, 9, 2, 3, 1, 5, 4, 2, 6}, then the function should return 3", func(t *testing.T) {
		a := []int{2, 1, 1, 2, 3, 1}
		expected := 3

		got := Solution(a)

		assert.Equal(t, expected, got)

	})
}
