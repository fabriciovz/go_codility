package number_solitaire

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberSolitaire(t *testing.T) {
	t.Run("Given A={1, -2, 0, 9, -1, -2}, the function should return 8", func(t *testing.T) {
		a := []int{1, -2, 0, 9, -1, -2}
		expected := 8

		got := Solution(a)

		assert.Equal(t, expected, got)
	})
}
