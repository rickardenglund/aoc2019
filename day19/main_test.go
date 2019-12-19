package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_part2Res(t *testing.T) {
	path = "input.txt"
	assert.Equal(t, 70_013, part2Solve(2))
}

func Test_part2Res2(t *testing.T) {
	path = "input.txt"
	assert.Equal(t, 140_026, part2Solve(3))
}
