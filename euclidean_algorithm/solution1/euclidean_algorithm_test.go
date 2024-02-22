package solution1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEuclideanAlgorithm(t *testing.T) {
	t.Run("Given A=50 and B=12, the function should return greatest common divisor GCD = 2", func(t *testing.T) {
		A := 50
		B := 12
		expected := 2

		got := Solution(A, B)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A=83 and B=19, the function should return greatest common divisor GCD = 2", func(t *testing.T) {
		A := 83
		B := 19
		expected := 1

		got := Solution(A, B)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A=15 and B=10, the function should return greatest common divisor GCD = 5", func(t *testing.T) {
		A := 15
		B := 10
		expected := 5

		got := Solution(A, B)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A=18 and B=12, the function should return greatest common divisor GCD = 5", func(t *testing.T) {
		A := 18
		B := 12
		expected := 6

		got := Solution(A, B)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A=9 and B=6, the function should return greatest common divisor GCD = 6", func(t *testing.T) {
		A := 9
		B := 6
		expected := 3

		got := Solution(A, B)

		assert.Equal(t, expected, got)
	})
}
