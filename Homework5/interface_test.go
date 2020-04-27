package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, figure Figure, want float64) {
		t.Helper()
		got := figure.Area()
		assert.Equal(t, got, want)
	}

	t.Run("squares", func(t *testing.T) {
		square := Square{5}
		checkArea(t, square, 25.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

}
