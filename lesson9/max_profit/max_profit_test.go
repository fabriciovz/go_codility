package max_profit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	t.Run("Given A := []int{4, 3, 4, 4,4,2}, the function should return 2", func(t *testing.T) {
		A := []int{23171, 21011, 21123, 21366, 21013, 21367}
		expected := 356

		got := Solution(A)

		assert.Equal(t, expected, got)
	})
}
