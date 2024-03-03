package sieve_of_eratosthenes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountFactor(t *testing.T) {
	{
		t.Run("Given N=24, the function should return {2, 3, 5, 7, 11, 13, 17, 19,23}", func(t *testing.T) {
			N := 24
			expected := []int{2, 3, 5, 7, 11, 13, 17, 19, 23}

			got := Solution(N)

			assert.Equal(t, expected, got)
		})
		t.Run("Given N=20, the function should return {2, 3, 5, 7, 11, 13, 17, 19}", func(t *testing.T) {
			N := 20
			expected := []int{2, 3, 5, 7, 11, 13, 17, 19}

			got := Solution(N)

			assert.Equal(t, expected, got)
		})
		t.Run("Given N=16, the function should return {2, 3, 5, 7, 11, 13}", func(t *testing.T) {
			N := 16
			expected := []int{2, 3, 5, 7, 11, 13}

			got := Solution(N)

			assert.Equal(t, expected, got)
		})
		t.Run("Given N=15, the function should return {2, 3, 5, 7, 11, 13}", func(t *testing.T) {
			N := 15
			expected := []int{2, 3, 5, 7, 11, 13}

			got := Solution(N)

			assert.Equal(t, expected, got)
		})
		t.Run("Given N=19, the function should return {2, 3, 5, 7, 11, 13, 17, 19}", func(t *testing.T) {
			N := 19
			expected := []int{2, 3, 5, 7, 11, 13, 17, 19}

			got := Solution(N)

			assert.Equal(t, expected, got)
		})
		t.Run("Given N=8, the function should return {2, 3, 5, 7}", func(t *testing.T) {
			N := 8
			expected := []int{2, 3, 5, 7}

			got := Solution(N)

			assert.Equal(t, expected, got)
		})
	}
}
