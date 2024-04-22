package equileader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEquileader(t *testing.T) {
	t.Run("Given A := []int{4, 3, 4, 4, 4, 2}, the function should return 2", func(t *testing.T) {
		A := []int{4, 3, 4, 4, 4, 2}
		expected := 2

		got := Solution(A)

		assert.Equal(t, expected, got)
	})
}
