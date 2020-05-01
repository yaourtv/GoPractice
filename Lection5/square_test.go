package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestEndEqual(t *testing.T) {

	var s = Square{Point{1,1},5}
	endVal1, endVal2 := s.End()
	testVal1, testVal2 := 6, 6
	assert.Equal(t, endVal1, testVal1, "X position values should be equal")
	assert.Equal(t, endVal2, testVal2, "Y position values should be equal")
}

func TestEndDifferent(t *testing.T) {

	var s = Square{Point{1,4},6}
	endVal1, endVal2 := s.End()
	testVal1, testVal2 := 7, 10
	assert.Equal(t, endVal1, testVal1, "X position values should be equal")
	assert.Equal(t, endVal2, testVal2, "Y position values should be equal")
}

func TestPerimeter(t *testing.T) {

	var s = Square{Point{1,3},5}
	assert.Equal(t, s.Perimeter(), uint(20), "perimeter should be equal length * 4")
}

func TestArea(t *testing.T) {

	var s = Square{Point{2,4},6}
	assert.Equal(t, s.Area(), uint(36), "area should be equal to the square of length")
}