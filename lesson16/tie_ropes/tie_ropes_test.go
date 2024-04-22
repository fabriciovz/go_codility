package tie_ropes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTieRopes(t *testing.T) {
	t.Run("Given A := []int{1, 2, 3, 4, 1, 1, 3} and K := 4, then the function should return 3", func(t *testing.T) {
		ropes := []int{1, 2, 3, 4, 1, 1, 3}
		K := 4
		expected := 3

		got := Solution(K, ropes)

		assert.Equal(t, expected, got)
	})
}
