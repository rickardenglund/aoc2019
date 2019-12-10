package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const ex1 = `.#..#
.....
#####
....#
...##`

func Test_getAsteroids(t *testing.T) {
	m := getAsteroids(ex1)
	//fmt.Printf("%v\n", m)
	assert.Equal(t, 10, len(m))
}

func Test_inLine1(t *testing.T) {
	res := inLine(
		pos{0, 0},
		pos{1, 1},
		pos{2, 2})
	assert.True(t, res)

	res = inLine(
		pos{0, 0},
		pos{1, 1},
		pos{-1, -1})
	assert.True(t, res)

	res = inLine(
		pos{0, 0},
		pos{1, 2},
		pos{-1, -1})
	assert.False(t, res)

	res = inLine(
		pos{4, 0},
		pos{2, 2},
		pos{1, 2})
	assert.False(t, res)

	res = inLine(
		pos{4, 0},
		pos{4, 2},
		pos{1, 0})
	assert.False(t, res)
}

func Test_canSee(t *testing.T) {
	m := getAsteroids(ex1)
	assert.True(t, canSee(pos{4, 0}, pos{2, 2}, m))
	assert.True(t, canSee(pos{4, 0}, pos{1, 2}, m))
	assert.True(t, canSee(pos{4, 0}, pos{1, 0}, m))
	assert.True(t, canSee(pos{1, 0}, pos{2, 2}, m))
	assert.False(t, canSee(pos{4, 2}, pos{2, 2}, m))
	assert.False(t, canSee(pos{4, 2}, pos{4, 4}, m))
	assert.True(t, canSee(pos{4, 2}, pos{4, 0}, m))
}

func Test_NAsteroids(t *testing.T) {
	m := getAsteroids(ex1)
	assert.Equal(t, 7, visibleAsteroids(pos{1, 0}, m))
	assert.Equal(t, 5, visibleAsteroids(pos{4, 2}, m))
	assert.Equal(t, 7, visibleAsteroids(pos{4, 0}, m))
}

func Test_NAsteroids2(t *testing.T) {
	m := getAsteroids(`......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`)
	assert.Equal(t, 33, visibleAsteroids(pos{5, 8}, m))
}

func Test_NAsteroidsLarge(t *testing.T) {
	m := getAsteroids(`.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`)
	//fmt.Printf("%v\n", m)
	assert.Equal(t, 210, visibleAsteroids(pos{11, 13}, m))
}
