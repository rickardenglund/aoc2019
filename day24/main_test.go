package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testState = `....#
#..#.
#..##
..#..
#....`

func Test_bioFunc(t *testing.T) {
	m := findRepeatingState(testState)
	assert.Equal(t, 2129920, bioRating(m))
}
