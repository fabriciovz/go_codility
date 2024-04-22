package solution1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountDiv(t *testing.T) {
	t.Run("Given A=6,B=11,K=2, the function should return 3", func(t *testing.T) {
		A := 6
		B := 11
		K := 2
		expected := 3

		got := Solution(A, B, K)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A=11,B=10,K=2, the function should return 0", func(t *testing.T) {
		A := 11
		B := 10
		K := 2
		expected := 0

		got := Solution(A, B, K)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A=0,B=0,K=11, the function should return 1", func(t *testing.T) {
		A := 0
		B := 0
		K := 11
		expected := 1

		got := Solution(A, B, K)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A=11,B=345,K=17, the function should return 20", func(t *testing.T) {
		A := 11
		B := 345
		K := 17
		expected := 20

		got := Solution(A, B, K)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A=11,B=11,K=11, the function should return 1", func(t *testing.T) {
		A := 11
		B := 11
		K := 11
		expected := 1

		got := Solution(A, B, K)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A=0,B=2000000000,K=1, the function should return 2000000001", func(t *testing.T) {
		A := 0
		B := 2000000000
		K := 1
		expected := 2000000001

		got := Solution(A, B, K)

		assert.Equal(t, expected, got)
	})
	t.Run("Given A=25,B=70,K=10, the function should return 5", func(t *testing.T) {
		A := 25
		B := 70
		K := 10
		expected := 5

		got := Solution(A, B, K)

		assert.Equal(t, expected, got)
	})
}
