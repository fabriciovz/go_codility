package cyclic_rotation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCyclicRotation(t *testing.T) {
	t.Run("Given A = [3, 8, 9, 7, 6] and K= -1, the function should return [3, 8, 9, 7, 6]", func(t *testing.T) {
		var a = []int{3, 8, 9, 7, 6}
		var k = -1
		var expected = []int{3, 8, 9, 7, 6}

		got := Solution(a, k)

		assert.Equal(t, expected, got)
	})

	t.Run("Given A = [3, 8, 9, 7, 6] and K= 101, the function should return [3, 8, 9, 7, 6]", func(t *testing.T) {
		var a = []int{3, 8, 9, 7, 6}
		var k = 101
		var expected = []int{3, 8, 9, 7, 6}

		got := Solution(a, k)

		assert.Equal(t, expected, got)
	})

	t.Run("Given A = [3, 8, 9, 7, 6] and K= 0, the function should return [3, 8, 9, 7, 6]", func(t *testing.T) {
		var a = []int{3, 8, 9, 7, 6}
		var k = 0
		var expected = []int{3, 8, 9, 7, 6}

		got := Solution(a, k)

		assert.Equal(t, expected, got)
	})

	t.Run("Given A = [3, 8, 9, 7, 6] and K= 1, the function should return [6, 3, 8, 9, 7]", func(t *testing.T) {
		var a = []int{3, 8, 9, 7, 6}
		var k = 1
		var expected = []int{6, 3, 8, 9, 7}

		got := Solution(a, k)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A = [3, 8, 9, 7, 6] and K= 2, the function should return [7, 6, 3, 8, 9]", func(t *testing.T) {
		var a = []int{3, 8, 9, 7, 6}
		var k = 2
		var expected = []int{7, 6, 3, 8, 9}

		got := Solution(a, k)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A = [3, 8, 9, 7, 6] and K= 3, the function should return [9, 7, 6, 3, 8]", func(t *testing.T) {
		var a = []int{3, 8, 9, 7, 6}
		var k = 3
		var expected = []int{9, 7, 6, 3, 8}

		got := Solution(a, k)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A = [0, 0, 0] and K= 4, the function should return [0, 0, 0]", func(t *testing.T) {
		var a = []int{0, 0, 0}
		var k = 4
		var expected = []int{0, 0, 0}

		got := Solution(a, k)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A = [1, 2, 3, 4] and K= 4, the function should return [1, 2, 3, 4]", func(t *testing.T) {
		var a = []int{1, 2, 3, 4}
		var k = 4
		var expected = []int{1, 2, 3, 4}

		got := Solution(a, k)

		assert.Equal(t, expected, got)
	})

}
