package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sequence(t *testing.T) {
	seq := `deal with increment 7
deal into new stack
deal into new stack`

	deckLength := 10
	deck := make([]int, deckLength)
	for i := 0; i < deckLength; i++ {
		res := executeFollow(i, deckLength, seq)
		deck[res] = i
	}
	assert.Equal(t, []int{0, 3, 6, 9, 2, 5, 8, 1, 4, 7}, deck)
}

func Test_sequence2(t *testing.T) {
	seq := `cut 6
deal with increment 7
deal into new stack`

	deckLength := 10
	deck := make([]int, deckLength)
	for i := 0; i < deckLength; i++ {
		res := executeFollow(i, deckLength, seq)
		deck[res] = i
	}
	assert.Equal(t, []int{3, 0, 7, 4, 1, 8, 5, 2, 9, 6}, deck)
}

func Test_sequence3(t *testing.T) {
	seq := `deal with increment 7
deal with increment 9
cut -2`

	deckLength := 10
	deck := make([]int, deckLength)
	for i := 0; i < deckLength; i++ {
		res := executeFollow(i, deckLength, seq)
		deck[res] = i
	}
	assert.Equal(t, []int{6, 3, 0, 7, 4, 1, 8, 5, 2, 9}, deck)
}

func Test_sequence4(t *testing.T) {
	seq := `deal into new stack
cut -2
deal with increment 7
cut 8
cut -4
deal with increment 7
cut 3
deal with increment 9
deal with increment 3
cut -1`

	deckLength := 10
	deck := make([]int, deckLength)
	for i := 0; i < deckLength; i++ {
		res := executeFollow(i, deckLength, seq)
		deck[res] = i
	}
	assert.Equal(t, []int{9, 2, 5, 8, 1, 4, 7, 0, 3, 6}, deck)
}
