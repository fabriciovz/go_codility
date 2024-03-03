package slices

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceTricks(t *testing.T) {
	t.Run("", func(t *testing.T) {
		A := []string{"Caicara", "Jusepín", "La Toscana", "Maturín"}
		expected := []string{"Caicara", "Jusepín", "La Toscana"}

		got := RemoveLastElement(A)

		assert.Equal(t, expected, got)
	})
	t.Run("", func(t *testing.T) {
		A := []string{"Caicara", "Jusepín", "La Toscana", "Maturín"}
		index := 2
		expected := []string{"Caicara", "Jusepín", "Maturín"}

		got := RemoveIndex(A, index)

		assert.Equal(t, expected, got)
	})
	t.Run("", func(t *testing.T) {
		A := []string{"Caicara", "Jusepín", "La Toscana", "Maturín"}
		index1 := 1
		index2 := 2
		expected := []string{"Caicara", "La Toscana", "Jusepín", "Maturín"}

		got := SwapElements(A, index1, index2)

		assert.Equal(t, expected, got)
	})
	t.Run("", func(t *testing.T) {
		A := []int{9, 1, 3, 20, 5}
		expected := []int{1, 3, 5, 9, 20}

		got := SortASC(A)

		assert.Equal(t, expected, got)
	})
	t.Run("", func(t *testing.T) {
		A := []int{9, 1, 3, 20, 5}
		expected := []int{20, 9, 5, 3, 1}

		got := SortDSC(A)

		assert.Equal(t, expected, got)
	})
}
