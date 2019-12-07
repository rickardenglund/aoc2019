package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalc(t *testing.T) {
	mem := []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}

	assert.Equal(t, 43210, CalcSignal([]int{4, 3, 2, 1, 0}, mem))
}

func TestCalc2(t *testing.T) {
	mem := []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}

	assert.Equal(t, 54321, CalcSignal([]int{0, 1, 2, 3, 4}, mem))
}

func TestGetSequence(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, getSequence(123))
}
