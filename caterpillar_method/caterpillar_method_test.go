package caterpillar_method

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	t.Run("Given N=0, then the function should return 0", func(t *testing.T) {
		N := []int{6, 2, 7, 4, 1, 3, 6}
		S := 12
		expected := true

		got := Solution(N, S)

		assert.Equal(t, expected, got)
	})
}
