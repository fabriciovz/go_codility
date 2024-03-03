package solution2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDominator(t *testing.T) {
	t.Run("Given A= {3,4,3,2,3,-1,3,3}", func(t *testing.T) {
		a := []int{3, 4, 3, 2, 3, -1, 3, 3}
		expected := []int{0, 2, 4, 6, 7}

		got := Solution(a)

		assert.Contains(t, expected, got)
	})
	t.Run("Given A= {6,8,4,6,8,6,6}", func(t *testing.T) {
		a := []int{6, 8, 4, 6, 8, 6, 6}
		expected := []int{0, 3, 5, 6}

		got := Solution(a)

		assert.Contains(t, expected, got)
	})
}
