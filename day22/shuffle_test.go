package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseFollow(t *testing.T) {
	deckLength := 10
	pos := 3

	newPos := dealIntoNewStackFollow(pos, deckLength)
	assert.Equal(t, 6, newPos)
}

func Test_CutPositive(t *testing.T) {
	deckLength := 10
	N := 3

	res := make([]int, deckLength)
	for i := 0; i < deckLength; i++ {
		newPos := cutFollow(i, N, deckLength)
		res[newPos] = i
	}

	assert.Equal(t, []int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2}, res)
}

func Test_CutNegative(t *testing.T) {
	deckLength := 10
	N := -4

	res := make([]int, deckLength)
	for i := 0; i < deckLength; i++ {
		newPos := cutFollow(i, N, deckLength)
		res[newPos] = i
	}

	assert.Equal(t, []int{6, 7, 8, 9, 0, 1, 2, 3, 4, 5}, res)
}

func Test_dealWithIncrement2(t *testing.T) {
	deckLength := 10
	N := 3

	res := make([]int, deckLength)
	for i := 0; i < deckLength; i++ {
		newPos := dealWithIncrementFollow(i, N, deckLength)
		res[newPos] = i
	}

	assert.Equal(t, []int{0, 7, 4, 1, 8, 5, 2, 9, 6, 3}, res)
}
