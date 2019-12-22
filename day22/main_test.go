package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_dealIntoNewStack(t *testing.T) {
	deck := getDeck(4)
	dealIntoNewStack(deck)
	assert.Equal(t, []int{3, 2, 1, 0}, deck)
}

func Test_cutN(t *testing.T) {
	deck := getDeck(5)
	cut(deck, 2)
	assert.Equal(t, []int{2, 3, 4, 0, 1}, deck)
}

func Test_cutNNegative(t *testing.T) {
	deck := getDeck(5)
	cut(deck, -2)
	assert.Equal(t, []int{3, 4, 0, 1, 2}, deck)
}

func Test_cutNNegative2(t *testing.T) {
	deck := getDeck(10)
	cut(deck, -2)
	assert.Equal(t, []int{8, 9, 0, 1, 2, 3, 4, 5, 6, 7}, deck)
}

func Test_dealWithIncrement(t *testing.T) {
	deck := getDeck(10)
	dealWithIncrement(deck, 3)
	assert.Equal(t, []int{0, 7, 4, 1, 8, 5, 2, 9, 6, 3}, deck)
}

func Test_sequence(t *testing.T) {
	seq := `deal with increment 7
deal into new stack
deal into new stack`

	deck := getDeck(10)
	execute(deck, seq)
	assert.Equal(t, []int{0, 3, 6, 9, 2, 5, 8, 1, 4, 7}, deck)
}

func Test_sequence2(t *testing.T) {
	seq := `cut 6
deal with increment 7
deal into new stack`

	deck := getDeck(10)
	execute(deck, seq)
	assert.Equal(t, []int{3, 0, 7, 4, 1, 8, 5, 2, 9, 6}, deck)
}

func Test_sequence3(t *testing.T) {
	seq := `deal with increment 7
deal with increment 9
cut -2`

	deck := getDeck(10)
	execute(deck, seq)
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

	deck := getDeck(10)
	execute(deck, seq)
	assert.Equal(t, []int{9, 2, 5, 8, 1, 4, 7, 0, 3, 6}, deck)
}
