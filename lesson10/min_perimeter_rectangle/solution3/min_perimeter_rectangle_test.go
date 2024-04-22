package solution3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinPerimeterRectangle(t *testing.T) {
	{
		t.Run("Given N=30, the function should return 22", func(t *testing.T) {
			N := 30
			expected := 22

			got := Solution(N)

			assert.Equal(t, expected, got)
		})
		t.Run("Given N=100000000, the function should return 40000", func(t *testing.T) {
			N := 100000000
			expected := 40000

			got := Solution(N)

			assert.Equal(t, expected, got)
		})
		t.Run("Given N=1, the function should return 4", func(t *testing.T) {
			N := 1
			expected := 4

			got := Solution(N)

			assert.Equal(t, expected, got)
		})
		t.Run("Given N=982451653, the function should return 1964903308", func(t *testing.T) {
			N := 982451653
			expected := 1964903308

			got := Solution(N)

			assert.Equal(t, expected, got)
		})
	}
}
