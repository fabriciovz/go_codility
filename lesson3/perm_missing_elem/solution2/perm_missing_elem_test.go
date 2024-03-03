package solution2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermMissingElem(t *testing.T) {

	t.Run("Given A = {2,3,1,5}, then the function should return 4", func(t *testing.T) {
		a := []int{2, 3, 1, 5}
		expected := 4

		got := Solution(a)

		assert.Equal(t, expected, got)

	})
	t.Run("Given A = {2}, then the function should return 1", func(t *testing.T) {
		a := []int{2}
		expected := 1

		got := Solution(a)

		assert.Equal(t, expected, got)

	})
	t.Run("Given A = {}, then the function should return 1", func(t *testing.T) {
		a := []int{}
		expected := 1

		got := Solution(a)

		assert.Equal(t, expected, got)

	})
	t.Run("Given A = {}, then the function should return 1", func(t *testing.T) {
		a := []int{1, 3}
		expected := 2

		got := Solution(a)

		assert.Equal(t, expected, got)

	})
	t.Run("Given A = {}, then the function should return 1", func(t *testing.T) {
		a := []int{3, 4, 1}
		expected := 2

		got := Solution(a)

		assert.Equal(t, expected, got)

	})
	t.Run("Given A = {1,2}, then the function should return 3", func(t *testing.T) {
		a := []int{1, 2}
		expected := 3

		got := Solution(a)

		assert.Equal(t, expected, got)

	})
}
