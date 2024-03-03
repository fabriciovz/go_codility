package solution1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountFactor(t *testing.T) {
	{
		t.Run("Given N=24, the function should return 8", func(t *testing.T) {
			N := 24
			expected := 8

			got := Solution(N)

			assert.Equal(t, expected, got)
		})
		t.Run("Given N=20, the function should return 6", func(t *testing.T) {
			N := 20
			expected := 6

			got := Solution(N)

			assert.Equal(t, expected, got)
		})
		t.Run("Given N=16, the function should return 5", func(t *testing.T) {
			N := 16
			expected := 5

			got := Solution(N)

			assert.Equal(t, expected, got)
		})
	}

}
