package solution2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFrogImpl(t *testing.T) {
	t.Run("Given X = 10, Y = 85, D = 30, the function should return 3", func(t *testing.T) {
		x := 10
		y := 85
		d := 30
		expected := 3

		got := Solution(x, y, d)

		assert.Equal(t, expected, got)
	})
	t.Run("Given X = 50, Y = 1000, D = 50, the function should return 19", func(t *testing.T) {
		x := 50
		y := 1000
		d := 50
		expected := 19

		got := Solution(x, y, d)

		assert.Equal(t, expected, got)
	})
	t.Run("Given X = 50, Y = 1000, D = 50, the function should return 19", func(t *testing.T) {
		x := 50
		y := 1000
		d := 50
		expected := 19

		got := Solution(x, y, d)

		assert.Equal(t, expected, got)
	})
	t.Run("Given X = 50, Y = 1000, D = 0, the function should return 0", func(t *testing.T) {
		x := 50
		y := 1000
		d := 0
		expected := 0

		got := Solution(x, y, d)

		assert.Equal(t, expected, got)
	})
	t.Run("Given X = 50, Y = 50, D = 10, the function should return 0", func(t *testing.T) {
		x := 50
		y := 50
		d := 0
		expected := 0

		got := Solution(x, y, d)

		assert.Equal(t, expected, got)
	})
	t.Run("Given X = 100, Y = 50, D = 10, the function should return 0", func(t *testing.T) {
		x := 100
		y := 50
		d := 0
		expected := 0

		got := Solution(x, y, d)

		assert.Equal(t, expected, got)
	})
	t.Run("Given X = 1, Y = 1, D = 3, the function should return 0", func(t *testing.T) {
		x := 1
		y := 1
		d := 3
		expected := 0

		got := Solution(x, y, d)

		assert.Equal(t, expected, got)
	})
	t.Run("Given X = 1, Y = 1, D = 3, the function should return 0", func(t *testing.T) {
		x := 5
		y := 1000000000
		d := 2
		expected := 499999998

		got := Solution(x, y, d)

		assert.Equal(t, expected, got)
	})
}
