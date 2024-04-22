package distinct

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDistinct(t *testing.T) {
	t.Run("Given A = {2, 1, 1, 2, 3, 1}, then the function should return 3", func(t *testing.T) {
		a := []int{2, 1, 1, 2, 3, 1}
		expected := 3

		got := Solution(a)

		assert.Equal(t, expected, got)

	})
}
