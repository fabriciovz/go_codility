package common_prime_divisors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommonPrimeDivisors(t *testing.T) {
	t.Run("Given A := []int{15, 10, 3} and B := []int{75, 30, 5}, then the function should return 1", func(t *testing.T) {
		A := []int{15, 10, 3}
		B := []int{75, 30, 5}
		expected := 1

		got := Solution(A, B)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A := []int{15, 4, 2} and B := []int{75, 6, 16}, then the function should return 2", func(t *testing.T) {
		A := []int{15, 4, 2}
		B := []int{75, 6, 16}
		expected := 2

		got := Solution(A, B)

		assert.Equal(t, expected, got)
	})
}
