package solution2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoinChange(t *testing.T) {
	t.Run("Given N=0, then the function should return 0", func(t *testing.T) {
		coins := []int{5, 2, 1}
		amount := 11
		expected := []int{5, 5, 1}

		got := Solution(coins, amount)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N=0, then the function should return 0", func(t *testing.T) {
		coins := []int{10, 2, 1}
		amount := 15
		expected := []int{10, 2, 2, 1}

		got := Solution(coins, amount)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N=0, then the function should return 0", func(t *testing.T) {
		coins := []int{10, 6, 1}
		amount := 13
		expected := []int{10, 1, 1, 1}
		//limitation: it should be: {6,6,1}

		got := Solution(coins, amount)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N=0, then the function should return 0", func(t *testing.T) {
		coins := []int{2}
		amount := 3
		expected := []int{}
		//limitation: when there is no change available

		got := Solution(coins, amount)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N=0, then the function should return 0", func(t *testing.T) {
		coins := []int{1, 5, 10}
		amount := 12
		expected := 4

		got := Solution(coins, amount)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N=0, then the function should return 0", func(t *testing.T) {
		coins := []int{1, 5, 10}
		amount := 8
		expected := 2

		got := Solution(coins, amount)

		assert.Equal(t, expected, got)
	})
	t.Run("Given N=0, then the function should return 0", func(t *testing.T) {
		coins := []int{1, 5, 10}
		amount := 10
		expected := 4

		got := Solution(coins, amount)

		assert.Equal(t, expected, got)
	})
}
