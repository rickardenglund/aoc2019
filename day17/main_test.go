package main

import (
	"aoc2019/position"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const example = `..#..........
..#..........
#######...###
#.#...#...#.#
#############
..#...#...#..
..#####...^..`

func Test_getAlignment(t *testing.T) {

	m := map[position.Pos]int{}
	rows := strings.Split(example, "\n")
	for y := range rows {
		for x := range rows[y] {
			m[position.Pos{X: x, Y: y}] = int(rows[y][x])
		}
	}
	assert.Equal(t, 76, getAlignment(m))
	draw(m, 0, 20, 0, 20)
}
