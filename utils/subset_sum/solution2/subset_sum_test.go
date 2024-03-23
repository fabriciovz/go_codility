package solution2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinAbsSum(t *testing.T) {
	t.Run("Given A := []int{1, 5, 2, 10} and Target := 7, then the function should return true", func(t *testing.T) {
		A := []int{1, 5, 2, 10}
		Target := 7
		N := len(A)
		expected := true

		got := Solution(A, N, Target)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{1, 5, 2, 10} and Target := 7, then the function should return true", func(t *testing.T) {
		A := []int{3, 34, 4, 12, 5, 2}
		Target := 9
		N := len(A)
		expected := true

		got := Solution(A, N, Target)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{2, 2, 3} and Target := 5, then the function should return true", func(t *testing.T) {
		A := []int{2, 2, 3}
		Target := 5
		N := len(A)
		expected := true

		got := Solution(A, N, Target)

		assert.Equal(t, expected, got)
	})
}
