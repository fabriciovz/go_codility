package solution1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCyclicRotation(t *testing.T) {
	t.Run("Given A = {[()()]} the function should return 1", func(t *testing.T) {
		var a = "{[()()]}"
		var expected = 1

		got := Solution(a)

		assert.Equal(t, expected, got)
	})

	t.Run("Given A = ([)()] , the function should return 0", func(t *testing.T) {
		var a = "([)()]"
		var expected = 0

		got := Solution(a)

		assert.Equal(t, expected, got)
	})

	t.Run("Given A = {[()()]} the function should return 1", func(t *testing.T) {
		var a = "(([]){}[])"
		var expected = 1

		got := Solution(a)

		assert.Equal(t, expected, got)
	})
}
