package go_coding

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryGap(t *testing.T) {
	t.Run("when 9 is pass, then it's got 2", func(t *testing.T) {
		binaryGap := Solution(9)

		assert.Equal(t, 2, binaryGap)
	})
	t.Run("when 529 is pass, then it's got 4", func(t *testing.T) {
		binary_gap := Solution(529)

		assert.Equal(t, 4, binary_gap)
	})
	t.Run("when 20 is pass, then it's got 1", func(t *testing.T) {
		binary_gap := Solution(20)

		assert.Equal(t, 1, binary_gap)
	})
	t.Run("when 15 is pass, then it's got 0", func(t *testing.T) {
		binary_gap := Solution(15)

		assert.Equal(t, 0, binary_gap)
	})
	t.Run("when 32 is pass, then it's got 0", func(t *testing.T) {
		binary_gap := Solution(32)

		assert.Equal(t, 0, binary_gap)
	})

}
