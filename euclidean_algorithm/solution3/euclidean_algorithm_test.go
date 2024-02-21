package solution3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEuclideanAlgorithm(t *testing.T) {
	t.Run("Given A=50 and B=12, the function should return greatest common divisor GDC = 2", func(t *testing.T) {
		A := 50
		B := 12
		expected := 2

		got := Solution(A, B, 1)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A=83 and B=19, the function should return greatest common divisor GDC = 2", func(t *testing.T) {
		A := 83
		B := 19
		expected := 1

		got := Solution(A, B, 1)

		assert.Equal(t, expected, got)
	})
}
