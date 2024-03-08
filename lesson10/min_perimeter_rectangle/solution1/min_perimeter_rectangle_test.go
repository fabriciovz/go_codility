package solution1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountFactor(t *testing.T) {
	{
		t.Run("Given N=24, the function should return 8", func(t *testing.T) {
			N := 30
			expected := 22

			got := Solution(N)

			assert.Equal(t, expected, got)
		})
		t.Run("Given N=24, the function should return 8", func(t *testing.T) {
			N := 100000000
			expected := 40000

			got := Solution(N)

			assert.Equal(t, expected, got)
		})
	}
}
